package controller

import (
	"github.com/gin-gonic/gin"
	firebase2 "github.com/knight7024/go-push-server/domain/firebase"
	"github.com/knight7024/go-push-server/domain/response"
	"github.com/knight7024/go-push-server/ent"
	"github.com/knight7024/go-push-server/server/handler"
)

// TopicSubscribe godoc
// @Summary		주제 구독
// @Description 여러 기기를 주제에 구독시킬 때 사용합니다.
// @Tags 		Firebase
// @Accept 		json
// @Produce 	json
// @Param 		X-Push-Client-Key	header 		string							true	"프로젝트의 `Client-Key`"
// @Param 		TopicWithTokens		body 		firebase.TopicWithTokens		true	"`TopicWithTokens` 예시"
// @Security	BearerAuth
// @Success 	200 				{object} 	response.FirebaseResponse
// @Failure 	400 				{object} 	response.errorResponse
// @Failure 	401 				{object} 	response.errorResponse
// @Failure 	403 				{object} 	response.errorResponse
// @Failure 	500 				{object} 	response.errorResponse
// @Router 		/api/topic/subscribe [post]
func TopicSubscribe(c *gin.Context) {
	p := c.MustGet("project").(*ent.Project)
	var req *firebase2.TopicWithTokens
	if err := c.ShouldBindJSON(&req); err != nil {
		ex := response.ErrorBuilder.NewWithError(response.BadRequestError).
			Reason(err.Error()).
			Build()
		c.JSON(ex.GetStatusCode(), ex)
		return
	}

	res := handler.TopicSubscribeHandler(p, req)
	c.JSON(res.GetStatusCode(), res)
}

// TopicUnsubscribe godoc
// @Summary		주제 구독 해제
// @Description 여러 기기를 구독 해제시킬 때 사용합니다.
// @Tags 		Firebase
// @Accept 		json
// @Produce 	json
// @Param 		X-Push-Client-Key	header 		string							true	"프로젝트의 `Client-Key`"
// @Param 		TopicWithTokens		body 		firebase.TopicWithTokens		true	"`TopicWithTokens` 예시"
// @Security	BearerAuth
// @Success 	200 				{object} 	response.FirebaseResponse
// @Failure 	400 				{object} 	response.errorResponse
// @Failure 	401 				{object} 	response.errorResponse
// @Failure 	403 				{object} 	response.errorResponse
// @Failure 	500 				{object} 	response.errorResponse
// @Router 		/api/topic/unsubscribe [post]
func TopicUnsubscribe(c *gin.Context) {
	p := c.MustGet("project").(*ent.Project)
	var req *firebase2.TopicWithTokens
	if err := c.ShouldBindJSON(&req); err != nil {
		ex := response.ErrorBuilder.NewWithError(response.BadRequestError).
			Reason(err.Error()).
			Build()
		c.JSON(ex.GetStatusCode(), ex)
		return
	}

	res := handler.TopicUnsubscribeHandler(p, req)
	c.JSON(res.GetStatusCode(), res)
}

// PushMessage godoc
// @Summary		푸시 알림 전송
// @Description 여러 개의 단일 메시지를 전송할 때 사용합니다.
// @Tags 		Firebase
// @Accept 		json
// @Produce 	json
// @Param 		X-Push-Client-Key	header 		string							true	"프로젝트의 `Client-Key`"
// @Param 		Messages			body 		firebase.Messages				true	"`Messages` 예시"
// @Security	BearerAuth
// @Success 	200 				{object} 	response.FirebaseResponse
// @Failure 	400 				{object} 	response.errorResponse
// @Failure 	401 				{object} 	response.errorResponse
// @Failure 	403 				{object} 	response.errorResponse
// @Failure 	500 				{object} 	response.errorResponse
// @Router 		/api/push/message [post]
func PushMessage(c *gin.Context) {
	p := c.MustGet("project").(*ent.Project)
	var req *firebase2.Messages
	if err := c.ShouldBindJSON(&req); err != nil {
		ex := response.ErrorBuilder.NewWithError(response.BadRequestError).
			Reason(err.Error()).
			Build()
		c.JSON(ex.GetStatusCode(), ex)
		return
	}

	res := handler.PushMessageHandler(p, req)
	c.JSON(res.GetStatusCode(), res)
}

// PushMulticast godoc
// @Summary		푸시 알림 전송
// @Description 하나의 메시지를 여러 기기에 전송할 때 사용합니다.
// @Tags 		Firebase
// @Accept 		json
// @Produce 	json
// @Param 		X-Push-Client-Key			header 		string								true	"프로젝트의 `Client-Key`"
// @Param 		CustomMulticastMessage		body 		firebase.CustomMulticastMessage		true	"`WithCustomMulticastMessage` 예시"
// @Security	BearerAuth
// @Success 	200 						{object} 	response.FirebaseResponse
// @Failure 	400 						{object} 	response.errorResponse
// @Failure 	401 						{object} 	response.errorResponse
// @Failure 	403 						{object} 	response.errorResponse
// @Failure 	500 						{object} 	response.errorResponse
// @Router 		/api/push/multicast [post]
func PushMulticast(c *gin.Context) {
	p := c.MustGet("project").(*ent.Project)
	var req *firebase2.CustomMulticastMessage
	if err := c.ShouldBindJSON(&req); err != nil {
		ex := response.ErrorBuilder.NewWithError(response.BadRequestError).
			Reason(err.Error()).
			Build()
		c.JSON(ex.GetStatusCode(), ex)
		return
	}

	res := handler.PushMulticastHandler(p, req)
	c.JSON(res.GetStatusCode(), res)
}
