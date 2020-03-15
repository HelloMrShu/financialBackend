package main

import (
	"github.com/astaxie/beego"
	_ "template/routers"
	"template/components"
)

func init() {
	components.InitLogger()
	components.InitDB()
}

func main() {
	beego.Run()
}