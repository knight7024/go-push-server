package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/knight7024/go-push-server/domain/response"
	"github.com/knight7024/go-push-server/domain/user"
	"regexp"
)

const (
	usernameRegExp = "^[a-zA-Z0-9]*.{4,32}$"
	passwordRegExp = "^[a-zA-Z0-9_!@#$_%^&*.?()-=+]*.{8,32}$"
)

func ValidateUser(c *gin.Context) {
	var req *user.User
	if err := c.ShouldBindJSON(&req); err != nil {
		ex := response.ErrorBuilder.NewWithError(response.BadRequestError).
			Reason(err.Error()).
			Build()
		c.AbortWithStatusJSON(ex.GetStatusCode(), ex)
		return
	}

	if u, _ := regexp.MatchString(usernameRegExp, req.Username); !u {
		ex := response.ErrorBuilder.NewWithError(response.UsernameValidationFailedError).
			Build()
		c.AbortWithStatusJSON(ex.GetStatusCode(), ex)
		return
	}
	if p, _ := regexp.MatchString(passwordRegExp, req.Password); !p {
		ex := response.ErrorBuilder.NewWithError(response.PasswordValidationFailedError).
			Build()
		c.AbortWithStatusJSON(ex.GetStatusCode(), ex)
		return
	}

	c.Set("user", req)
	c.Next()
}
