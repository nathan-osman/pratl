package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"

	_ "github.com/nathan-osman/pratl/models"
	_ "github.com/nathan-osman/pratl/routers"

	"fmt"
)

func init() {
	orm.RegisterDataBase(
		"default",
		"postgres",
		fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s",
			beego.AppConfig.String("db_user"),
			beego.AppConfig.String("db_pass"),
			beego.AppConfig.DefaultString("db_host", "127.0.0.1"),
			beego.AppConfig.DefaultString("db_port", "5432"),
			beego.AppConfig.String("db_name"),
		),
	)
}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
