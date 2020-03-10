package routers

import (
	"template/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/ae/test", &controllers.RuleController{}, "get:AeTest")
	beego.Router("/ex/test", &controllers.RuleController{}, "get:ExTest")
	beego.Router("/ex/prod", &controllers.RuleController{}, "get:ExProd")
}
