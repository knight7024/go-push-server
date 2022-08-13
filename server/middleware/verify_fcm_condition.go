package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/knight7024/go-push-server/domain/project"
	"github.com/knight7024/go-push-server/domain/response"
	"github.com/knight7024/go-push-server/ent"
)

func VerifyFCMCondition(c *gin.Context) {
	clientKey := c.GetHeader("X-Push-Client-Key")
	if _, err := uuid.Parse(clientKey); err != nil {
		ex := response.ErrorBuilder.NewWithError(response.InvalidClientKeyError).
			Build()
		c.AbortWithStatusJSON(ex.GetStatusCode(), ex)
		return
	}

	ret, err := project.ReadOneByClientKey(context.TODO(), clientKey)
	if ent.IsNotFound(err) {
		ex := response.ErrorBuilder.NewWithError(response.InvalidClientKeyError).
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

	c.Set("project", ret)
	c.Next()
}
