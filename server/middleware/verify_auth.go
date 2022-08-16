package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/knight7024/go-push-server/common/util"
	"github.com/knight7024/go-push-server/domain/response"
	"github.com/knight7024/go-push-server/domain/user"
	"github.com/knight7024/go-push-server/ent"
	"strings"
)

func VerifyAuthAndUser(c *gin.Context) {
	tempUID, ex := verifyAuth(c)
	if ex != nil {
		c.AbortWithStatusJSON(ex.GetStatusCode(), ex)
		return
	} else if tempUID == nil {
		ex = response.ErrorBuilder.NewWithError(response.InvalidTokenError).
			Build()
		c.AbortWithStatusJSON(ex.GetStatusCode(), ex)
	}

	if uid, ex := verifyUser(tempUID); ex != nil {
		c.AbortWithStatusJSON(ex.GetStatusCode(), ex)
	} else {
		c.Set("uid", uid)
	}

	c.Next()
}

func VerifyAuth(c *gin.Context) {
	if uid, ex := verifyAuth(c); ex != nil {
		c.AbortWithStatusJSON(ex.GetStatusCode(), ex)
		return
	} else if uid == nil {
		ex = response.ErrorBuilder.NewWithError(response.InvalidTokenError).
			Build()
		c.AbortWithStatusJSON(ex.GetStatusCode(), ex)
	}
	c.Next()
}

func verifyAuth(c *gin.Context) (any, response.Response) {
	header := strings.TrimSpace(c.GetHeader("Authorization"))
	if !strings.HasPrefix(header, "Bearer ") {
		return nil, response.ErrorBuilder.NewWithError(response.AuthenticationRequiredError).
			Build()
	}
	header = strings.TrimPrefix(header, "Bearer ")
	payload, err := util.Validate(util.AccessToken(header))
	if err != nil {
		if vErr, ok := err.(*jwt.ValidationError); ok {
			switch {
			case errors.As(vErr, &jwt.ErrTokenExpired):
				return nil, response.ErrorBuilder.NewWithError(response.ExpiredTokenError).
					Reason(vErr.Error()).
					Build()
			case errors.As(vErr, &jwt.ErrTokenUsedBeforeIssued):
				return nil, response.ErrorBuilder.NewWithError(response.BeforeIssuedTokenError).
					Reason(vErr.Error()).
					Build()
			}
		}
		return nil, response.ErrorBuilder.NewWithError(response.InvalidTokenError).
			Build()
	}
	return payload.Get("uid"), nil
}

func verifyUser(tempUID any) (uid int, ex response.Response) {
	switch uidType := tempUID.(type) {
	case json.Number:
		t, _ := uidType.Int64()
		uid = int(t)
	case float64:
		uid = int(uidType)
	default:
		return 0, response.ErrorBuilder.NewWithError(response.InvalidTokenError).
			Build()
	}

	_, err := user.ReadOneByUserID(context.TODO(), uid)
	if ent.IsNotFound(err) {
		return 0, response.ErrorBuilder.NewWithError(response.UserNotFoundError).
			Build()
	} else if err != nil {
		return 0, response.ErrorBuilder.NewWithError(response.DatabaseServerError).
			Reason(err.Error()).
			Build()
	}

	return
}
