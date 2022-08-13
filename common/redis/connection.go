package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/knight7024/go-push-server/common/config"
)

type ConnectionPool struct {
	*redis.Client
}

var Connection = new(ConnectionPool)

func (cp *ConnectionPool) InitConnection() {
	Connection.Client = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Address,
		Username: config.Config.Redis.Username,
		Password: config.Config.Redis.Password,
		DB:       config.Config.Redis.DBName,
	})
}
