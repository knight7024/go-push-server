package firebase

import (
	"firebase.google.com/go/messaging"
	"strings"
)

type Messages struct {
	Messages []*messaging.Message `json:"messages" binding:"required"`
}

// CustomMulticastMessage 는 다음과 같은 규칙을 따른다.
// condition이 존재할 경우, tokens와 topic 항목은 무시된다.
// condition이 존재하지 않고 topic이 존재하면, tokens 항목은 무시된다.
type CustomMulticastMessage struct {
	Tokens       []string                 `json:"tokens,omitempty"`
	Topic        string                   `json:"topic,omitempty"`
	Condition    string                   `json:"condition,omitempty"`
	Data         map[string]string        `json:"data,omitempty"`
	Notification *messaging.Notification  `json:"notification,omitempty" binding:"required"`
	Android      *messaging.AndroidConfig `json:"android,omitempty"`
	Webpush      *messaging.WebpushConfig `json:"webpush,omitempty"`
	APNS         *messaging.APNSConfig    `json:"apns,omitempty"`
}

func (cmm *CustomMulticastMessage) ToMessages() (messages []*messaging.Message) {
	if strings.TrimSpace(cmm.Condition) != "" {
		m := &messaging.Message{
			Data:         cmm.Data,
			Notification: cmm.Notification,
			Android:      cmm.Android,
			Webpush:      cmm.Webpush,
			APNS:         cmm.APNS,
			Condition:    cmm.Condition,
		}
		messages = append(messages, m)
	} else if strings.TrimSpace(cmm.Topic) != "" {
		m := &messaging.Message{
			Data:         cmm.Data,
			Notification: cmm.Notification,
			Android:      cmm.Android,
			Webpush:      cmm.Webpush,
			APNS:         cmm.APNS,
			Topic:        cmm.Topic,
		}
		messages = append(messages, m)
	} else {
		for _, token := range cmm.Tokens {
			m := &messaging.Message{
				Data:         cmm.Data,
				Notification: cmm.Notification,
				Android:      cmm.Android,
				Webpush:      cmm.Webpush,
				APNS:         cmm.APNS,
				Token:        token,
			}
			messages = append(messages, m)
		}
	}
	return
}
