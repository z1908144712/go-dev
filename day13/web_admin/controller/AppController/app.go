package AppController

import (
	"fmt"
	"go_dev/day13/web_admin/model"
	"strings"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

type AppController struct {
	beego.Controller
}

func (p *AppController) AppList() {
	p.Layout = "layout/layout.html"

	appList, err := model.GetAllAppInfo()
	if err != nil {
		logs.Error(err)
		p.Data["Error"] = fmt.Sprintf("服务器繁忙")
		p.TplName = "app/error.html"
		return
	}
	p.Data["appList"] = appList
	p.TplName = "app/index.html"
}

func (p *AppController) AppApply() {
	p.Layout = "layout/layout.html"
	p.TplName = "app/apply.html"
}

func (p *AppController) AppCreate() {

	appName := p.GetString("app_name")
	appType := p.GetString("app_type")
	developPath := p.GetString("develop_path")
	ipstr := p.GetString("iplist")

	p.Layout = "layout/layout.html"

	if len(appName) == 0 || len(appType) == 0 || len(developPath) == 0 || len(ipstr) == 0 {
		p.Data["Error"] = fmt.Sprintf("非法参数")
		p.TplName = "app/error.html"

		return
	}

	appInfo := &model.AppInfo{}
	appInfo.AppName = appName
	appInfo.AppType = appType
	appInfo.DevelopPath = developPath
	appInfo.IP = strings.Split(ipstr, ",")
	err := model.CreateApp(appInfo)
	if err != nil {
		p.Data["Error"] = fmt.Sprintf("创建项目失败，数据库繁忙")
		p.TplName = "app/error.html"

		logs.Error(err)
		return
	}

	p.Redirect("/app/list", 302)
}
