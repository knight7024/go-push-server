package handler

import (
	"context"
	"github.com/knight7024/go-push-server/common/util"
	"github.com/knight7024/go-push-server/domain/firebase"
	"github.com/knight7024/go-push-server/domain/response"
	"github.com/knight7024/go-push-server/ent"
	"net/http"
)

func TopicSubscribeHandler(p *ent.Project, req *firebase.TopicWithTokens) response.Response {
	firebaseUtil, err := util.InitApp(p.Credentials, p.ProjectID, p.ID)
	if err != nil {
		return response.ErrorBuilder.NewWithError(response.InternalFirebaseError).
			Reason(err.Error()).
			Build()
	}
	res, err := firebaseUtil.Subscribe(context.TODO(), req.Tokens, req.Topic)
	if err != nil {
		return response.ErrorBuilder.NewWithError(response.InternalFirebaseError).
			Reason(err.Error()).
			Build()
	}

	return response.SuccessBuilder.New(http.StatusOK).
		Data(&response.FirebaseResponse{
			SuccessCount: res.SuccessCount,
			FailureCount: res.FailureCount,
		}).
		Build()
}

func TopicUnsubscribeHandler(p *ent.Project, req *firebase.TopicWithTokens) response.Response {
	firebaseUtil, err := util.InitApp(p.Credentials, p.ProjectID, p.ID)
	if err != nil {
		return response.ErrorBuilder.NewWithError(response.InternalFirebaseError).
			Reason(err.Error()).
			Build()
	}
	res, err := firebaseUtil.Unsubscribe(context.TODO(), req.Tokens, req.Topic)
	if err != nil {
		return response.ErrorBuilder.NewWithError(response.InternalFirebaseError).
			Reason(err.Error()).
			Build()
	}

	return response.SuccessBuilder.New(http.StatusOK).
		Data(&response.FirebaseResponse{
			SuccessCount: res.SuccessCount,
			FailureCount: res.FailureCount,
		}).
		Build()
}

func PushMessageHandler(p *ent.Project, req *firebase.Messages) response.Response {
	firebaseUtil, err := util.InitApp(p.Credentials, p.ProjectID, p.ID)
	if err != nil {
		return response.ErrorBuilder.NewWithError(response.InternalFirebaseError).
			Reason(err.Error()).
			Build()
	}
	res, err := firebaseUtil.SendMessage(context.TODO(), req.Messages)
	if err != nil {
		return response.ErrorBuilder.NewWithError(response.InternalFirebaseError).
			Reason(err.Error()).
			Build()
	}

	return response.SuccessBuilder.New(http.StatusOK).
		Data(&response.FirebaseResponse{
			SuccessCount: res.SuccessCount,
			FailureCount: res.FailureCount,
		}).
		Build()
}

func PushMulticastHandler(p *ent.Project, req *firebase.CustomMulticastMessage) response.Response {
	firebaseUtil, err := util.InitApp(p.Credentials, p.ProjectID, p.ID)
	if err != nil {
		return response.ErrorBuilder.NewWithError(response.InternalFirebaseError).
			Reason(err.Error()).
			Build()
	}
	res, err := firebaseUtil.SendMessage(context.TODO(), req.ToMessages())
	if err != nil {
		return response.ErrorBuilder.NewWithError(response.InternalFirebaseError).
			Reason(err.Error()).
			Build()
	}

	return response.SuccessBuilder.New(http.StatusOK).
		Data(&response.FirebaseResponse{
			SuccessCount: res.SuccessCount,
			FailureCount: res.FailureCount,
		}).
		Build()
}
