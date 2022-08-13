package handler

import (
	"context"
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
		setCmd := rdb.Connection.Set(context.TODO(), fmt.Sprintf(refreshTokenCacheKey, ret.ID), refreshToken, util.RefreshTokenDuration)
		if setCmd.Err() != nil {
			return response.ErrorBuilder.NewWithError(response.RedisServerError).
				Reason(err.Error()).
				Build()
		}

		return response.SuccessBuilder.New(http.StatusOK).
			Data(&response.AuthTokens{
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
			}).
			Build()
	} else if err != nil {
		return response.ErrorBuilder.NewWithError(response.RedisServerError).
			Reason(err.Error()).
			Build()
	}

	return response.SuccessBuilder.New(http.StatusOK).
		Data(&response.AuthTokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
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

	accessToken, _ := util.AccessTokenBuilder.New().
		Set("uid", ret.ID).
		Build()
	refreshToken, _ := util.RefreshTokenBuilder.New().
		Set("uid", ret.ID).
		Build()
	setCmd := rdb.Connection.Set(context.TODO(), fmt.Sprintf(refreshTokenCacheKey, ret.ID), refreshToken, util.RefreshTokenDuration)
	if setCmd.Err() != nil {
		return response.ErrorBuilder.NewWithError(response.RedisServerError).
			Reason(err.Error()).
			Build()
	}

	return response.SuccessBuilder.New(http.StatusOK).
		Data(&response.AuthTokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}).
		Build()
}
