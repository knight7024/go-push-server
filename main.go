package main

import (
	"log"

	"github.com/knight7024/go-push-server/common/config"
	"github.com/knight7024/go-push-server/common/mysql"
	"github.com/knight7024/go-push-server/common/redis"
	_ "github.com/knight7024/go-push-server/docs"
	"github.com/knight7024/go-push-server/server"
	"github.com/spf13/viper"
)

func init() {
	setViperConfig()
	mysql.Connection.InitConnection()
	redis.Connection.InitConnection()
	config.InitCache()
}

func setViperConfig() {
	viper.SetConfigName("application_config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		switch err := err.(type) {
		case viper.ConfigFileNotFoundError:
			log.Fatalf("config file not found: %v\n", err)
		case viper.ConfigParseError:
			log.Fatalf("failed parsing config file: %v\n", err)
		default:
			log.Fatalf("fatal error config file: %v\n", err)
		}
	}
	if err := viper.Unmarshal(&config.Config); err != nil {
		log.Fatalf("fatal error config file: %v\n", err)
	}
}

// @title           			Push Server API
// @version         			1.0.0
// @description    				Push Server developed by Jongwoo Jeong

// @license.name  				MIT License
// @license.url   				https://opensource.org/licenses/MIT

// @securityDefinitions.apikey 	BearerAuth
// @in 							header
// @name 						Authorization
func main() {
	// MySQL 커넥션풀 종료
	defer mysql.Connection.Close()

	r := server.InitRouter()
	_ = r.Run()
}
