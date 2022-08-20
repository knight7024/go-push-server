package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/knight7024/go-push-server/server/controller"
	"github.com/knight7024/go-push-server/server/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func InitRouter() *gin.Engine {
	// 기본 엔진으로 gin 설정
	// Reverse Proxy 통해서 Request 들어오므로 localhost만 신뢰
	router := gin.Default()
	_ = router.SetTrustedProxies([]string{"127.0.0.1"})

	// Swagger URL Mapping
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health Check URL Mapping
	router.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// API URL Mapping
	apiURL := router.Group("/api")
	apiURL.Use(middleware.CORS)
	{
		verifyAuth := apiURL.Group("")
		verifyAuth.Use(middleware.VerifyAuthAndUser)
		{
			verifyAuth.GET("/project/all", controller.ReadAllProjects)
			verifyAuth.POST("/project", controller.CreateProject)
			verifyAuth.GET("/user/logout", controller.Logout)

			verifyFCMCondition := verifyAuth.Group("")
			verifyFCMCondition.Use(middleware.VerifyFCMCondition)
			{
				verifyFCMCondition.POST("/topic/subscribe", controller.TopicSubscribe)
				verifyFCMCondition.POST("/topic/unsubscribe", controller.TopicUnsubscribe)
				verifyFCMCondition.POST("/push/message", controller.PushMessage)
				verifyFCMCondition.POST("/push/multicast", controller.PushMulticast)
			}

			validateProjectID := verifyAuth.Group(fmt.Sprintf("%s/:id", "/project"))
			validateProjectID.Use(middleware.InputProjectID)
			{
				validateProjectID.GET("", controller.ReadProject)
				validateProjectID.PUT("", controller.UpdateProject)
				validateProjectID.DELETE("", controller.DeleteProject)
			}
		}

		apiURL.POST("/user/refresh", controller.RefreshTokens)

		validateUser := apiURL.Group("/user")
		validateUser.Use(middleware.ValidateUser)
		{
			validateUser.POST("/login", controller.Login)
			validateUser.POST("/signup", controller.Signup)
		}
	}

	return router
}
