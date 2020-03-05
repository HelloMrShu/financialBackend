package routers

import (
	"template/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/ae/test", &controllers.RuleController{}, "get:AeTest")
}
