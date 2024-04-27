package pgsql

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

func init() {
	pgsqlUser, _ := beego.AppConfig.String("pgsqlUser")
	pgsqlPass, _ := beego.AppConfig.String("pgsqlPass")
	pgsqlDbName, _ := beego.AppConfig.String("pgsqlDbName")
	pgsqlHost, _ := beego.AppConfig.String("pgsqlHost")
	pgsqlPort, _ := beego.AppConfig.String("pgsqlPort")
	pgsqlSslMode, _ := beego.AppConfig.String("pgsqlSslMode")
	// 注册驱动
	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		panic(err)
	}

	// 设置默认数据库
	err = orm.RegisterDataBase("default", "postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", pgsqlUser, pgsqlPass, pgsqlDbName, pgsqlHost, pgsqlPort, pgsqlSslMode))
	if err != nil {
		panic(err)
	}
}
func NewPgsqlOrm() orm.Ormer {
	ormer := orm.NewOrm()
	return ormer
}
