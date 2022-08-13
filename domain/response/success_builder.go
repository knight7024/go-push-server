package response

type successBuilder struct {
	statusCode int
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

func (rb *successBuilder) Data(data interface{}) *successBuilder {
	rb.data = data
	return rb
}

func (rb *successBuilder) Build() *successResponse {
	return &successResponse{
		StatusCode: rb.statusCode,
		Data:       rb.data,
	}
}
