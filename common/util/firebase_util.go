package util

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"github.com/knight7024/go-push-server/common/config"
	"google.golang.org/api/option"
)

type firebaseUtil struct {
	firebaseApp *firebase.App
	fcmClient   *messaging.Client
	pid         int
	dirty       bool
}

const ProjectCacheKey = "project:%d"

func (fu *firebaseUtil) MakeDirty() {
	if cached, exists := config.MemCache.Get(fmt.Sprintf(ProjectCacheKey, fu.pid)); exists {
		fu := cached.(*firebaseUtil)
		fu.dirty = true
		go updateFirebaseUtilCache(fu.pid, fu)
	}
}

func GetFirebaseUtilFromCache(pid int) (*firebaseUtil, bool) {
	if cached, exists := config.MemCache.Get(fmt.Sprintf(ProjectCacheKey, pid)); exists {
		return cached.(*firebaseUtil), true
	}
	return nil, false
}

func updateFirebaseUtilCache(pid int, fu *firebaseUtil) {
	config.MemCache.Set(fmt.Sprintf(ProjectCacheKey, pid), fu, 0)
}

func InitApp(credentials []byte, projectID string, pid int) (fu *firebaseUtil, err error) {
	var exists bool
	if fu, exists = GetFirebaseUtilFromCache(pid); !exists || fu.dirty {
		fu = &firebaseUtil{pid: pid}
		opt := option.WithCredentialsJSON(credentials)
		if fu.firebaseApp, err = firebase.NewApp(context.TODO(), &firebase.Config{ProjectID: projectID}, opt); err != nil {
			return nil, fmt.Errorf("error on initializing app: %v\n", err)
		}
		if fu.fcmClient, err = fu.firebaseApp.Messaging(context.TODO()); err != nil {
			return nil, fmt.Errorf("error on initializing messaging: %v", err)
		}
		fu.dirty = false
		go updateFirebaseUtilCache(pid, fu)
	}
	return
}

func (fu *firebaseUtil) Subscribe(ctx context.Context, tokens []string, topic string) (*messaging.TopicManagementResponse, error) {
	return fu.fcmClient.SubscribeToTopic(ctx, tokens, topic)
}

func (fu *firebaseUtil) Unsubscribe(ctx context.Context, tokens []string, topic string) (*messaging.TopicManagementResponse, error) {
	return fu.fcmClient.UnsubscribeFromTopic(ctx, tokens, topic)
}

func (fu *firebaseUtil) SendMessage(ctx context.Context, messages []*messaging.Message) (*messaging.BatchResponse, error) {
	return fu.fcmClient.SendAll(ctx, messages)
}
