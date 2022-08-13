package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/knight7024/go-push-server/domain/user"
	"github.com/knight7024/go-push-server/server/handler"
)

// Login godoc
// @Summary		로그인
// @Description 유저가 로그인할 때 사용합니다.
// @Tags 		User
// @Accept 		json
// @Produce 	json
// @Param 		User	body 		user.User					true	"User 예시"
// @Success 	200		{object}	response.AuthTokens
// @Failure 	400 	{object} 	response.errorResponse
// @Failure 	401 	{object} 	response.errorResponse
// @Failure 	500		{object} 	response.AuthTokens
// @Router 		/api/user/login [post]
func Login(c *gin.Context) {
	u, _ := c.Get("user")
	req := u.(*user.User)

	res := handler.LoginHandler(req)
	c.JSON(res.GetStatusCode(), res)
}

// Logout godoc
// @Summary		로그아웃
// @Description 유저가 로그아웃할 때 사용합니다.
// @Tags 		User
// @Produce 	json
// @Security	BearerAuth
// @Success 	204
// @Failure 	401 	{object} 	response.errorResponse
// @Failure 	500		{object} 	response.errorResponse
// @Router 		/api/user/logout [get]
func Logout(c *gin.Context) {
	uid := c.GetInt("uid")

	res := handler.LogoutHandler(uid)
	c.JSON(res.GetStatusCode(), res)
}

// Signup godoc
// @Summary		회원가입
// @Description 새로운 유저가 가입할 때 사용합니다.
// @Tags 		User
// @Accept 		json
// @Produce 	json
// @Param 		User	body 		user.User					true	"User 예시"
// @Success 	201		{object}	response.AuthTokens
// @Failure 	400 	{object} 	response.errorResponse
// @Failure 	409 	{object} 	response.errorResponse
// @Failure 	500		{object} 	response.errorResponse
// @Router 		/api/user/signup [post]
func Signup(c *gin.Context) {
	u, _ := c.Get("user")
	req := u.(*user.User)

	res := handler.SignupHandler(req)
	c.JSON(res.GetStatusCode(), res)
}
