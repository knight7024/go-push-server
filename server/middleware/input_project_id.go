package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/knight7024/go-push-server/domain/response"
)

func InputProjectID(c *gin.Context) {
	var pid int
	_, err := fmt.Sscan(c.Param("id"), &pid)
	if err != nil {
		ex := response.ErrorBuilder.NewWithError(response.BadParametersError).
			Reason(err.Error()).
			Build()
		c.AbortWithStatusJSON(ex.GetStatusCode(), ex)
		return
	}

	c.Set("pid", pid)
	c.Next()
}
