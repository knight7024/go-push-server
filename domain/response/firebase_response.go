package response

type FirebaseResponse struct {
	SuccessCount int         `json:"success_count"`
	FailureCount int         `json:"failure_count"`
	Data         interface{} `json:"data,omitempty" swaggertype:"object,object"`
}
