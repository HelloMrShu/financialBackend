package main

import (
	_ "template/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
    orm.Debug = true
    orm.RegisterDataBase("default", "mysql", "dbuser:dbuser@tcp(10.19.37.10:3306)/berry?charset=utf8", 30)
s}

func main() {
	beego.Run()
}