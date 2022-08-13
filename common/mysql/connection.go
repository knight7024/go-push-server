package mysql

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/knight7024/go-push-server/common/config"
	"github.com/knight7024/go-push-server/ent"
	"log"
)

type ConnectionPool struct {
	*ent.Client
}

var Connection = new(ConnectionPool)

func (cp *ConnectionPool) InitConnection() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4,utf8&parseTime=True&loc=Asia%%2FSeoul",
		config.Config.Datasource.Username,
		config.Config.Datasource.Password,
		config.Config.Datasource.Address,
		config.Config.Datasource.DBName,
	)
	var err error
	Connection.Client, err = ent.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	// Run the auto migration tool.
	if err = Connection.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
