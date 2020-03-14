package main

import (
	_ "template/routers"
	"github.com/astaxie/beego"
	"template/components"
)

func init() {
	components.InitDB()
	components.InitLogger()
}

func main() {
	beego.Run()
}