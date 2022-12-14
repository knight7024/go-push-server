package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/knight7024/go-push-server/common/config"
	"time"
)

const (
	AccessTokenDuration  = time.Minute * 30
	RefreshTokenDuration = time.Hour * 24 * 30
)

type Token string
type AccessToken string
type RefreshToken string

type JWT interface {
	Token | AccessToken | RefreshToken
	String() string
}

func (t Token) String() string {
	return string(t)
}

func (t AccessToken) String() string {
	return string(t)
}

func (t RefreshToken) String() string {
	return string(t)
}

func Validate(token string) (*claimsBuilder, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Config.Core.SecretKey), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return nil, err
	}
	return extractPayload(parsedToken)
}

func extractPayload(token *jwt.Token) (*claimsBuilder, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if !claims.VerifyIssuer(config.Config.Core.AppName, true) {
			return nil, errors.New("invalid token")
		}
		return &claimsBuilder{claims}, nil
	}
	return nil, errors.New("failed parsing payload")
}

type claimsBuilder struct {
	claims jwt.MapClaims
}

type tokenBuilder struct {
	Token
}

type specialClaimsBuilder struct {
	claims jwt.MapClaims
}

type accessTokenBuilder struct {
	AccessToken
}

type refreshTokenBuilder struct {
	RefreshToken
}

var TokenBuilder tokenBuilder
var AccessTokenBuilder accessTokenBuilder
var RefreshTokenBuilder refreshTokenBuilder

func (tokenBuilder) New() *claimsBuilder {
	return (&claimsBuilder{make(map[string]interface{})}).
		issuer(config.Config.Core.AppName)
}

func (accessTokenBuilder) New() *specialClaimsBuilder {
	builder := &specialClaimsBuilder{make(map[string]interface{})}
	builder.claims["iss"] = config.Config.Core.AppName
	builder.claims["exp"] = &jwt.NumericDate{Time: time.Now().Add(AccessTokenDuration)}
	return builder
}

func (refreshTokenBuilder) New() *specialClaimsBuilder {
	builder := &specialClaimsBuilder{make(map[string]interface{})}
	builder.claims["iss"] = config.Config.Core.AppName
	builder.claims["exp"] = &jwt.NumericDate{Time: time.Now().Add(RefreshTokenDuration)}
	return builder
}

func (cb *claimsBuilder) issuer(iss string) *claimsBuilder {
	cb.claims["iss"] = iss
	return cb
}

func (cb *claimsBuilder) Subject(sub string) *claimsBuilder {
	cb.claims["sub"] = sub
	return cb
}

func (cb *claimsBuilder) Audience(aud ...string) *claimsBuilder {
	cb.claims["aud"] = aud
	return cb
}

func (cb *claimsBuilder) ExpiresAt(exp time.Time) *claimsBuilder {
	cb.claims["exp"] = &jwt.NumericDate{Time: exp}
	return cb
}

func (cb *claimsBuilder) NotBefore(nbf time.Time) *claimsBuilder {
	cb.claims["nbf"] = &jwt.NumericDate{Time: nbf}
	return cb
}

func (cb *claimsBuilder) IssuedAt(iat time.Time) *claimsBuilder {
	cb.claims["iat"] = &jwt.NumericDate{Time: iat}
	return cb
}

func (cb *claimsBuilder) ID(jti string) *claimsBuilder {
	cb.claims["jti"] = jti
	return cb
}

func (cb *claimsBuilder) Get(key string) any {
	return cb.claims[key]
}

func (cb *claimsBuilder) Set(key string, value any) *claimsBuilder {
	cb.claims[key] = value
	return cb
}

func (cb *specialClaimsBuilder) Subject(sub string) *specialClaimsBuilder {
	cb.claims["sub"] = sub
	return cb
}

func (cb *specialClaimsBuilder) Audience(aud ...string) *specialClaimsBuilder {
	cb.claims["aud"] = aud
	return cb
}

func (cb *specialClaimsBuilder) NotBefore(nbf time.Time) *specialClaimsBuilder {
	cb.claims["nbf"] = &jwt.NumericDate{Time: nbf}
	return cb
}

func (cb *specialClaimsBuilder) IssuedAt(iat time.Time) *specialClaimsBuilder {
	cb.claims["iat"] = &jwt.NumericDate{Time: iat}
	return cb
}

func (cb *specialClaimsBuilder) ID(jti string) *specialClaimsBuilder {
	cb.claims["jti"] = jti
	return cb
}

func (cb *specialClaimsBuilder) Get(key string) any {
	return cb.claims[key]
}

func (cb *specialClaimsBuilder) Set(key string, value any) *specialClaimsBuilder {
	cb.claims[key] = value
	return cb
}

func (cb *claimsBuilder) Build() (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, cb.claims).
		SignedString([]byte(config.Config.Core.SecretKey))
}

func (cb *specialClaimsBuilder) Build() (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, cb.claims).
		SignedString([]byte(config.Config.Core.SecretKey))
}
