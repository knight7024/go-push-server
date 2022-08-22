package middleware

import (
	"github.com/gin-gonic/gin"
)

func CORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin, X-Push-Client-Key")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS")
	c.Next()
}
