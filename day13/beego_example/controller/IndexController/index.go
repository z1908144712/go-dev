package IndexController

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (p *IndexController) Index() {
	//p.TplName = "index.html"
	m := make(map[string]interface{})
	m["name"] = "name"
	m["age"] = 18
	p.Data["json"] = m
	p.ServeJSON()
}
