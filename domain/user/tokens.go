package user

type Tokens interface {
	GetString() string
}

type AccessToken struct {
	AccessToken string `json:"access_token,omitempty"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token,omitempty"`
}

type AuthTokens struct {
	*AccessToken
	*RefreshToken
}
