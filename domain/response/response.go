package response

import (
	"github.com/goccy/go-json"
)

type Response interface {
	GetStatusCode() int
}

type successResponse struct {
	StatusCode int         `json:"-"`
	Message    string      `json:"success_message,omitempty"`
	Data       interface{} `json:"data,omitempty" swaggertype:"object,object"`
}

func (s *successResponse) MarshalJSON() ([]byte, error) {
	if s.Data != nil {
		type rawMessage struct {
			json.RawMessage
		}
		data, _ := json.Marshal(s.Data)
		return json.Marshal(&struct {
			Message     string `json:"success_message,omitempty"`
			*rawMessage `json:",omitempty"`
		}{
			Message:    s.Message,
			rawMessage: &rawMessage{data},
		})
	}
	return json.Marshal(&struct {
		Message string `json:"success_message,omitempty"`
	}{
		Message: s.Message,
	})
}

func (s *successResponse) GetStatusCode() int {
	return s.StatusCode
}

type errorResponse struct {
	Code       string      `json:"error_code"`
	StatusCode int         `json:"-"`
	Message    string      `json:"error_message"`
	Reason     string      `json:"reason,omitempty"`
	Data       interface{} `json:"data,omitempty" swaggertype:"object,object"`
}

func (e *errorResponse) GetStatusCode() int {
	return e.StatusCode
}
