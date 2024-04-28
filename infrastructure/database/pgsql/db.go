package pgsql

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"zhiqu/infrastructure/config"
)

func NewPgsqlDB() *gorm.DB {
	var PostgresConfig = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", config.AppConfig.Pgsql.Host, config.AppConfig.Pgsql.User, config.AppConfig.Pgsql.Password, config.AppConfig.Pgsql.Database, config.AppConfig.Pgsql.Port, config.AppConfig.Pgsql.SslMode)
	// 连接数据库
	db, err := gorm.Open(postgres.Open(PostgresConfig))
	if err != nil {
		panic(err)
	}
	return db
}
