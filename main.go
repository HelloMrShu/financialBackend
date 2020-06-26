package main

import (
	"github.com/astaxie/beego"
	_ "funds/routers"
	"funds/components"
)

func init() {
	components.InitLogger()
	components.InitDB()
}

func main() {
	beego.Run()
}