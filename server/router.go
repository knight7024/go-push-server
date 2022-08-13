package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/knight7024/go-push-server/common/config"
	"github.com/knight7024/go-push-server/server/controller"
	"github.com/knight7024/go-push-server/server/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	// 기본 엔진으로 gin 설정
	// Reverse Proxy 통해서 Request 들어오므로 localhost만 신뢰
	router := gin.Default()
	_ = router.SetTrustedProxies([]string{"127.0.0.1"})

	// Swagger URL Mapping
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API URL Mapping
	apiURL := router.Group("/api")
	apiURL.Use(middleware.CORS)
	{
		verifyAuth := apiURL.Group("")
		verifyAuth.Use(middleware.VerifyAuth)
		{
			verifyAuth.GET(config.Config.APIs.ProjectAllURI, controller.ReadAllProjects)
			verifyAuth.POST(config.Config.APIs.ProjectURI, controller.CreateProject)
			verifyAuth.GET(config.Config.APIs.UserLogoutURI, controller.Logout)

			verifyFCMCondition := verifyAuth.Group("")
			verifyFCMCondition.Use(middleware.VerifyFCMCondition)
			{
				verifyFCMCondition.POST(config.Config.APIs.TopicSubscribeURI, controller.TopicSubscribe)
				verifyFCMCondition.POST(config.Config.APIs.TopicUnsubscribeURI, controller.TopicUnsubscribe)
				verifyFCMCondition.POST(config.Config.APIs.PushMessageURI, controller.PushMessage)
				verifyFCMCondition.POST(config.Config.APIs.PushMulticastURI, controller.PushMulticast)
			}

			validateProjectID := verifyAuth.Group(fmt.Sprintf("%s/:id", config.Config.APIs.ProjectURI))
			validateProjectID.Use(middleware.InputProjectID)
			{
				validateProjectID.GET("", controller.ReadProject)
				validateProjectID.PUT("", controller.UpdateProject)
				validateProjectID.DELETE("", controller.DeleteProject)
			}
		}

		validateUser := apiURL.Group("")
		validateUser.Use(middleware.ValidateUser)
		{
			validateUser.POST(config.Config.APIs.UserLoginURI, controller.Login)
			validateUser.POST(config.Config.APIs.UserSignupURI, controller.Signup)
		}
	}

	return router
}
