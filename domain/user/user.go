package user

import (
	"context"
	"github.com/knight7024/go-push-server/common/mysql"
	"github.com/knight7024/go-push-server/ent"
	predicate "github.com/knight7024/go-push-server/ent/user"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id,omitempty" swaggerignore:"true"`
	Username string `json:"username,omitempty" binding:"required" minLength:"4" maxLength:"32"`
	Password string `json:"password,omitempty" binding:"required" minLength:"8" maxLength:"32"`
}

func ReadOneByUserID(ctx context.Context, uid int) (*ent.User, error) {
	return mysql.Connection.User.Query().
		Where(predicate.And(predicate.ID(uid), predicate.IsApproved(true))).
		Only(ctx)
}

func ReadOneByUsername(ctx context.Context, username string) (*ent.User, error) {
	return mysql.Connection.User.Query().
		Where(predicate.Username(username)).
		Only(ctx)
}

func CreateOne(ctx context.Context, user *User) (*ent.User, error) {
	encrypted, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	return mysql.Connection.User.Create().
		SetUsername(user.Username).
		SetPassword(string(encrypted)).
		Save(ctx)
}

func CountByUsername(ctx context.Context, username string) (int, error) {
	return mysql.Connection.User.Query().
		Where(predicate.Username(username)).
		Count(ctx)
}
