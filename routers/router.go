package routers

import (
	"funds/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/sector/list", &controllers.SectorController{}, "get:SectorList")
	beego.Router("/sector/save", &controllers.SectorController{}, "post:SectorSave")
	beego.Router("/sector/delete", &controllers.SectorController{}, "post:SectorDelete")
}
