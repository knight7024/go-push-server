package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/knight7024/go-push-server/domain/project"
	"github.com/knight7024/go-push-server/domain/response"
	"github.com/knight7024/go-push-server/server/handler"
)

// ReadAllProjects godoc
// @Summary		모든 프로젝트 읽기
// @Description 유저의 모든 프로젝트를 불러올 때 사용합니다.
// @Tags 		Project
// @Accept 		json
// @Produce 	json
// @Security	BearerAuth
// @Success 	200 			{object} 	[]ent.Project
// @Failure 	401 			{object} 	response.errorResponse
// @Failure 	500 			{object} 	response.errorResponse
// @Router 		/api/project/all [get]
func ReadAllProjects(c *gin.Context) {
	uid := c.GetInt("uid")

	res := handler.ReadAllProjectsHandler(uid)
	c.JSON(res.GetStatusCode(), res)
}

// CreateProject godoc
// @Summary		프로젝트 생성
// @Description 프로젝트를 생성할 때 사용합니다.
// @Tags 		Project
// @Accept 		json
// @Produce 	json
// @Param 		Project				body 				project.Project			true					"`Project` 예시"
// @Security	BearerAuth
// @Success 	201 				{object} 			ent.Project
// @Failure 	400 				{object} 			response.errorResponse
// @Failure 	401 				{object} 			response.errorResponse
// @Failure 	409 				{object} 			response.errorResponse
// @Failure 	500 				{object} 			response.errorResponse
// @Router 		/api/project [post]
func CreateProject(c *gin.Context) {
	uid := c.GetInt("uid")
	var req *project.Project
	if err := c.ShouldBindJSON(&req); err != nil {
		ex := response.ErrorBuilder.NewWithError(response.BadRequestError).
			Reason(err.Error()).
			Build()
		c.JSON(ex.GetStatusCode(), ex)
		return
	}

	res := handler.CreateProjectHandler(uid, req)
	c.JSON(res.GetStatusCode(), res)
}

// ReadProject godoc
// @Summary		단일 프로젝트 읽기
// @Description 유저의 프로젝트 하나를 불러올 때 사용합니다.
// @Tags 		Project
// @Accept 		json
// @Produce 	json
// @Param 		project_id		path 		int						true	"Project ID"
// @Security	BearerAuth
// @Success 	200 			{object} 	ent.Project
// @Failure 	401 			{object} 	response.errorResponse
// @Failure 	404 			{object} 	response.errorResponse
// @Failure 	500 			{object} 	response.errorResponse
// @Router 		/api/project/{project_id} [get]
func ReadProject(c *gin.Context) {
	uid, pid := c.GetInt("uid"), c.GetInt("pid")

	res := handler.ReadProjectHandler(uid, pid)
	c.JSON(res.GetStatusCode(), res)
}

// UpdateProject godoc
// @Summary		단일 프로젝트 수정
// @Description 유저의 프로젝트 하나를 수정할 때 사용합니다.
// @Tags 		Project
// @Accept 		json
// @Produce 	json
// @Param 		project_id		path 		int							true	"Project ID"
// @Param 		Project			body 		project.Project				true	"`Project` 예시"
// @Security	BearerAuth
// @Success 	200 			{object} 	ent.Project
// @Failure 	401 			{object} 	response.errorResponse
// @Failure 	404 			{object} 	response.errorResponse
// @Failure 	500 			{object} 	response.errorResponse
// @Router 		/api/project/{project_id} [patch]
func UpdateProject(c *gin.Context) {
	uid, pid := c.GetInt("uid"), c.GetInt("pid")
	var req *project.Project
	if err := c.ShouldBindJSON(&req); err != nil {
		ex := response.ErrorBuilder.NewWithError(response.BadRequestError).
			Reason(err.Error()).
			Build()
		c.JSON(ex.GetStatusCode(), ex)
		return
	}

	res := handler.UpdateProjectHandler(uid, pid, req)
	c.JSON(res.GetStatusCode(), res)
}

// UpdateProjectClientKey godoc
// @Summary		프로젝트의 Client Key 갱신
// @Description 프로젝트의 Client Key를 새롭게 갱신할 때 사용합니다.
// @Tags 		Project
// @Accept 		json
// @Produce 	json
// @Param 		project_id		path 		int							true	"Project ID"
// @Security	BearerAuth
// @Success 	200 			{object} 	response.OnlyClientKey
// @Failure 	401 			{object} 	response.errorResponse
// @Failure 	404 			{object} 	response.errorResponse
// @Failure 	500 			{object} 	response.errorResponse
// @Router 		/api/project/{project_id}/client-key [patch]
func UpdateProjectClientKey(c *gin.Context) {
	uid, pid := c.GetInt("uid"), c.GetInt("pid")

	res := handler.UpdateProjectClientKeyHandler(uid, pid)
	c.JSON(res.GetStatusCode(), res)
}

// DeleteProject godoc
// @Summary		단일 프로젝트 삭제
// @Description 유저의 프로젝트 하나를 삭제할 때 사용합니다.
// @Tags 		Project
// @Accept 		json
// @Produce 	json
// @Param 		project_id		path 		int							true	"Project ID"
// @Security	BearerAuth
// @Success 	204
// @Failure 	401 			{object} 	response.errorResponse
// @Failure 	404 			{object} 	response.errorResponse
// @Failure 	500 			{object} 	response.errorResponse
// @Router 		/api/project/{project_id} [delete]
func DeleteProject(c *gin.Context) {
	pid := c.GetInt("pid")

	res := handler.DeleteProjectHandler(pid)
	c.JSON(res.GetStatusCode(), res)
}
