package response

type successBuilder struct {
	statusCode int
	message    string
	data       interface{}
}

type sb struct{}

var SuccessBuilder *sb

func (b *sb) New(statusCode int) *successBuilder {
	return &successBuilder{statusCode: statusCode}
}

func (rb *successBuilder) StatusCode(statusCode int) *successBuilder {
	rb.statusCode = statusCode
	return rb
}

func (rb *successBuilder) Message(message string) *successBuilder {
	rb.message = message
	return rb
}

func (rb *successBuilder) Data(data interface{}) *successBuilder {
	rb.data = data
	return rb
}

func (rb *successBuilder) Build() *successResponse {
	return &successResponse{
		StatusCode: rb.statusCode,
		Message:    rb.message,
		Data:       rb.data,
	}
}
