package main

import (
	_ "firstbee/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // 引入mysql的驱动包
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "rich:www.123.com@tcp(192.168.120.78:3306)/blog?charset=utf8mb4")

}

func main() {
	beego.Run()
}
