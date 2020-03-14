package main

import (
	_ "template/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", "dbuser:dbuser@tcp(10.19.37.10:3306)/berry?charset=utf8", 30)
	orm.RegisterDataBase("test1", "mysql", "dbuser:dbuser@tcp(10.19.37.10:3306)/exchange_bata?charset=utf8", 30)
}

func main() {
	beego.Run()
}