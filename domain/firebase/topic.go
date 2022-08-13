package firebase

type TopicWithTokens struct {
	Topic  string   `json:"topic" binding:"required"`
	Tokens []string `json:"tokens" binding:"required"`
}
