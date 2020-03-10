package main

import (
	_ "template/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.Debug = true
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
}

func main() {
	beego.Run()
}