package LogController

import (
	"fmt"
	"go_dev/day13/web_admin/model"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

type LogController struct {
	beego.Controller
}

func (p *LogController) LogList() {
	p.Layout = "layout/layout.html"

	logList, err := model.GetAllLogInfo()
	if err != nil {
		logs.Error(err)
		p.Data["Error"] = fmt.Sprintf("服务器繁忙")
		p.TplName = "log/error.html"
		return
	}
	p.Data["logList"] = logList
	p.TplName = "log/index.html"
}

func (p *LogController) LogApply() {
	p.Layout = "layout/layout.html"
	p.TplName = "log/apply.html"
}

func (p *LogController) LogCreate() {

	appName := p.GetString("app_name")
	logPath := p.GetString("log_path")
	topic := p.GetString("topic")

	p.Layout = "layout/layout.html"

	if len(appName) == 0 || len(logPath) == 0 || len(topic) == 0 {
		p.Data["Error"] = fmt.Sprintf("非法参数")
		p.TplName = "log/error.html"

		return
	}

	logInfo := &model.LogInfo{}
	logInfo.AppName = appName
	logInfo.Topic = topic
	logInfo.LogPath = logPath
	err := model.CreateLog(logInfo)
	if err != nil {
		p.Data["Error"] = fmt.Sprintf("创建项目失败，数据库繁忙")
		p.TplName = "log/error.html"

		logs.Error(err)
		return
	}
	iplist, err := model.GetIPInfoByName(appName)
	if err != nil {
		p.Data["Error"] = fmt.Sprintf("获取项目失败，数据库繁忙")
		p.TplName = "log/error.html"

		logs.Error(err)
		return
	}
	for _, ip := range iplist {
		key := "/logadmin/backend/config" + ip
		err = model.SetEtcdConf(key, logInfo)
		if err != nil {
			logs.Error(err)
			continue
		}
	}

	p.Redirect("/log/list", 302)
}
