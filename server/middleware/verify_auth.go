package middleware

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/knight7024/go-push-server/common/util"
	"github.com/knight7024/go-push-server/domain/response"
	"github.com/knight7024/go-push-server/domain/user"
	"github.com/knight7024/go-push-server/ent"
	"strings"
)

func VerifyAuth(c *gin.Context) {
	header := strings.TrimSpace(c.GetHeader("Authorization"))
	if !strings.HasPrefix(header, "Bearer ") {
		ex := response.ErrorBuilder.NewWithError(response.AuthenticationRequiredError).
			Build()
		c.AbortWithStatusJSON(ex.GetStatusCode(), ex)
		return
	}

	header = strings.TrimPrefix(header, "Bearer ")
	payload, err := util.Validate(header)
	if err != nil {
		if vErr, ok := err.(jwt.ValidationError); ok {
			switch vErr.Unwrap() {
			case jwt.ErrTokenExpired:
				ex := response.ErrorBuilder.NewWithError(response.ExpiredTokenError).
					Reason(vErr.Error()).
					Build()
				c.AbortWithStatusJSON(ex.GetStatusCode(), ex)
				return
			case jwt.ErrTokenUsedBeforeIssued:
				ex := response.ErrorBuilder.NewWithError(response.BeforeIssuedTokenError).
					Reason(vErr.Error()).
					Build()
				c.AbortWithStatusJSON(ex.GetStatusCode(), ex)
				return
			}
		}
		ex := response.ErrorBuilder.NewWithError(response.InvalidTokenError).
			Reason(err.Error()).
			Build()
		c.AbortWithStatusJSON(ex.GetStatusCode(), ex)
		return
	}

	var uid int
	switch uidType := payload.Get("uid").(type) {
	case json.Number:
		t, _ := uidType.Int64()
		uid = int(t)
	case float64:
		uid = int(uidType)
	default:
		ex := response.ErrorBuilder.NewWithError(response.InvalidTokenError).
			Build()
		c.AbortWithStatusJSON(ex.GetStatusCode(), ex)
		return
	}
	_, err = user.ReadOneByUserID(context.TODO(), uid)
	if ent.IsNotFound(err) {
		ex := response.ErrorBuilder.NewWithError(response.UserNotFoundError).
			Build()
		c.AbortWithStatusJSON(ex.GetStatusCode(), ex)
		return
	} else if err != nil {
		ex := response.ErrorBuilder.NewWithError(response.DatabaseServerError).
			Reason(err.Error()).
			Build()
		c.AbortWithStatusJSON(ex.GetStatusCode(), ex)
		return
	}

	c.Set("uid", uid)
	c.Next()
}
