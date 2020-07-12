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

	beego.Router("/fund/list", &controllers.FundController{}, "get:FundList")
	beego.Router("/fund/save", &controllers.FundController{}, "post:FundSave")
	beego.Router("/fund/delete", &controllers.FundController{}, "post:FundDelete")
	beego.Router("/fund/update", &controllers.FundController{}, "post:FundUpdate")
}
