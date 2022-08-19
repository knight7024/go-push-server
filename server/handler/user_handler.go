package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	rdb "github.com/knight7024/go-push-server/common/redis"
	"github.com/knight7024/go-push-server/common/util"
	"github.com/knight7024/go-push-server/domain/response"
	"github.com/knight7024/go-push-server/domain/user"
	"github.com/knight7024/go-push-server/ent"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

const (
	refreshTokenCacheKey = "user:%d:refresh_token"
)

func LoginHandler(req *user.User) response.Response {
	ret, err := user.ReadOneByUsername(context.TODO(), req.Username)
	if ent.IsNotFound(err) {
		return response.UserNotFoundError
	} else if err != nil {
		return response.ErrorBuilder.NewWithError(response.DatabaseServerError).
			Reason(err.Error()).
			Build()
	}

	// 비밀번호 검증
	err = bcrypt.CompareHashAndPassword([]byte(ret.Password), []byte(req.Password))
	if err != nil {
		return response.PasswordNotMatchedError
	}

	accessToken, _ := util.AccessTokenBuilder.New().
		Set("uid", ret.ID).
		Build()
	refreshToken, err := rdb.Connection.Get(context.TODO(), fmt.Sprintf(refreshTokenCacheKey, ret.ID)).Result()
	if err == redis.Nil {
		refreshToken, _ = util.RefreshTokenBuilder.New().
			Set("uid", ret.ID).
			Build()
		cmdErr := rdb.Connection.Set(context.TODO(), fmt.Sprintf(refreshTokenCacheKey, ret.ID), refreshToken, util.RefreshTokenDuration).Err()
		if cmdErr != nil {
			return response.ErrorBuilder.NewWithError(response.RedisServerError).
				Reason(cmdErr.Error()).
				Build()
		}

		return response.SuccessBuilder.New(http.StatusOK).
			Data(&user.AuthTokens{
				AccessToken:  &user.AccessToken{AccessToken: accessToken},
				RefreshToken: &user.RefreshToken{RefreshToken: refreshToken},
			}).
			Build()
	} else if err != nil {
		return response.ErrorBuilder.NewWithError(response.RedisServerError).
			Reason(err.Error()).
			Build()
	}

	return response.SuccessBuilder.New(http.StatusOK).
		Data(&user.AuthTokens{
			AccessToken:  &user.AccessToken{AccessToken: accessToken},
			RefreshToken: &user.RefreshToken{RefreshToken: refreshToken},
		}).
		Build()
}

func LogoutHandler(uid int) response.Response {
	rdb.Connection.Del(context.TODO(), fmt.Sprintf(refreshTokenCacheKey, uid))

	return response.SuccessBuilder.New(http.StatusNoContent).
		Build()
}

func SignupHandler(req *user.User) response.Response {
	if count, err := user.CountByUsername(context.TODO(), req.Username); count != 0 {
		return response.AlreadyExistsUsernameError
	} else if err != nil {
		return response.ErrorBuilder.NewWithError(response.DatabaseServerError).
			Reason(err.Error()).
			Build()
	}

	ret, err := user.CreateOne(context.TODO(), req)
	if err != nil {
		return response.ErrorBuilder.NewWithError(response.DatabaseServerError).
			Reason(err.Error()).
			Build()
	}

	refreshToken, _ := util.RefreshTokenBuilder.New().
		Set("uid", ret.ID).
		Build()
	cmdErr := rdb.Connection.Set(context.TODO(), fmt.Sprintf(refreshTokenCacheKey, ret.ID), refreshToken, util.RefreshTokenDuration).Err()
	if cmdErr != nil {
		return response.ErrorBuilder.NewWithError(response.RedisServerError).
			Reason(cmdErr.Error()).
			Build()
	}

	return response.SuccessBuilder.New(http.StatusCreated).
		Message("가입이 완료되었습니다. 승인을 대기해주세요.").
		Build()
}

func RefreshTokensHandler(req *user.RefreshToken) response.Response {
	payload, err := util.Validate(util.RefreshToken(req.RefreshToken))
	if err != nil {
		return response.ErrorBuilder.NewWithError(response.InvalidTokenError).
			Build()
	}

	var uid int
	switch uidType := payload.Get("uid").(type) {
	case json.Number:
		t, _ := uidType.Int64()
		uid = int(t)
	case float64:
		uid = int(uidType)
	default:
		return response.ErrorBuilder.NewWithError(response.InvalidTokenError).
			Build()
	}
	if result, err := rdb.Connection.Get(context.TODO(), fmt.Sprintf(refreshTokenCacheKey, uid)).Result(); result != "" && result != req.RefreshToken {
		return response.ErrorBuilder.NewWithError(response.InvalidTokenError).
			Build()
	} else if err != nil {
		return response.ErrorBuilder.NewWithError(response.RedisServerError).
			Reason(err.Error()).
			Build()
	}

	accessToken, _ := util.AccessTokenBuilder.New().
		Set("uid", uid).
		Build()
	refreshToken, _ := util.RefreshTokenBuilder.New().
		Set("uid", uid).
		Build()
	cmdErr := rdb.Connection.Set(context.TODO(), fmt.Sprintf(refreshTokenCacheKey, uid), refreshToken, util.RefreshTokenDuration).Err()
	if cmdErr != nil {
		return response.ErrorBuilder.NewWithError(response.RedisServerError).
			Reason(cmdErr.Error()).
			Build()
	}

	return response.SuccessBuilder.New(http.StatusOK).
		Data(&user.AuthTokens{
			AccessToken:  &user.AccessToken{AccessToken: accessToken},
			RefreshToken: &user.RefreshToken{RefreshToken: refreshToken},
		}).
		Build()
}
