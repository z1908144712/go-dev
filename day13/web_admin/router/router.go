package router

import (
	"go_dev/day13/web_admin/controller/AppController"
	"go_dev/day13/web_admin/controller/LogController"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &AppController.AppController{}, "*:AppList")
	beego.Router("/app/list", &AppController.AppController{}, "*:AppList")
	beego.Router("/app/apply", &AppController.AppController{}, "*:AppApply")
	beego.Router("/app/create", &AppController.AppController{}, "*:AppCreate")

	beego.Router("/log/list", &LogController.LogController{}, "*:LogList")
	beego.Router("/log/apply", &LogController.LogController{}, "*:LogApply")
	beego.Router("/log/create", &LogController.LogController{}, "*:LogCreate")
}
