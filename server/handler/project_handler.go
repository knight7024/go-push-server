package handler

import (
	"context"
	"github.com/knight7024/go-push-server/common/util"
	"github.com/knight7024/go-push-server/domain/project"
	"github.com/knight7024/go-push-server/domain/response"
	"github.com/knight7024/go-push-server/ent"
	"net/http"
)

func ReadAllProjectsHandler(uid int) response.Response {
	ret, err := project.ReadAllByUserID(context.TODO(), uid)
	if err != nil {
		return response.ErrorBuilder.NewWithError(response.DatabaseServerError).
			Reason(err.Error()).
			Build()
	}

	return response.SuccessBuilder.New(http.StatusOK).
		Data(ret).
		Build()
}

func CreateProjectHandler(uid int, req *project.Project) response.Response {
	if count, err := project.CountByProjectID(context.TODO(), req.ProjectID); err != nil {
		return response.ErrorBuilder.NewWithError(response.DatabaseServerError).
			Reason(err.Error()).
			Build()
	} else if count != 0 {
		return response.AlreadyExistsProjectIDError
	}

	ret, err := project.CreateOne(context.TODO(), uid, req)
	if ent.IsConstraintError(err) {
		return response.AlreadyExistsProjectIDError
	} else if err != nil {
		return response.ErrorBuilder.NewWithError(response.DatabaseServerError).
			Reason(err.Error()).
			Build()
	}

	return response.SuccessBuilder.New(http.StatusCreated).
		Data(&response.OnlyID{
			ID: ret.ID,
		}).
		Build()
}

func ReadProjectHandler(uid int, pid int) response.Response {
	ret, err := project.ReadOneByUserID(context.TODO(), uid, pid)
	if ent.IsNotFound(err) {
		return response.DataNotFoundError
	} else if err != nil {
		return response.ErrorBuilder.NewWithError(response.DatabaseServerError).
			Reason(err.Error()).
			Build()
	}

	return response.SuccessBuilder.New(http.StatusOK).
		Data(ret).
		Build()
}

func UpdateProjectHandler(uid int, pid int, req *project.Project) response.Response {
	err := project.UpdateOneByUserID(context.TODO(), uid, pid, req)
	if ent.IsNotFound(err) {
		return response.DataNotFoundError
	} else if err != nil {
		return response.ErrorBuilder.NewWithError(response.DatabaseServerError).
			Reason(err.Error()).
			Build()
	}
	go func(p int) {
		if firebaseUtil, exists := util.GetFirebaseUtilFromCache(p); exists {
			firebaseUtil.MakeDirty()
		}
	}(pid)

	return response.SuccessBuilder.New(http.StatusNoContent).
		Build()
}

func UpdateProjectClientKeyHandler(uid int, pid int) response.Response {
	clientKey, err := project.UpdateClientKeyByUserID(context.TODO(), uid, pid)
	if ent.IsNotFound(err) {
		return response.DataNotFoundError
	} else if err != nil {
		return response.ErrorBuilder.NewWithError(response.DatabaseServerError).
			Reason(err.Error()).
			Build()
	}

	return response.SuccessBuilder.New(http.StatusOK).
		Data(&response.OnlyClientKey{ClientKey: clientKey}).
		Build()
}

func DeleteProjectHandler(pid int) response.Response {
	err := project.DeleteOneByUserID(context.TODO(), pid)
	if ent.IsNotFound(err) {
		return response.DataNotFoundError
	} else if err != nil {
		return response.ErrorBuilder.NewWithError(response.DatabaseServerError).
			Reason(err.Error()).
			Build()
	}

	return response.SuccessBuilder.New(http.StatusNoContent).
		Build()
}
