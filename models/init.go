package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dburl := beego.AppConfig.String("dburl")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbName := beego.AppConfig.String("db")
	//注册mysql Driver
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//构造conn连接
	conn := dbuser + ":" + dbpassword + "@tcp(" + dburl + ")/" + dbName + "?charset=utf8"
	//注册数据库连接
	err := orm.RegisterDataBase("default", "mysql", conn)
	if err != nil {
		panic(err)
	}
	fmt.Println("database connect success !")
}
