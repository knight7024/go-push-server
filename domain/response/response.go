package response

import "github.com/goccy/go-json"

type Response interface {
	GetStatusCode() int
}

type successResponse struct {
	StatusCode int         `json:"-"`
	Data       interface{} `json:"data,omitempty" swaggertype:"object,object"`
}

func (s *successResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Data)
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
