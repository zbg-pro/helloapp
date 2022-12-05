package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

// @Title Get
// @Description Get
// @Param
// @Success 200
// @Failure 403 "default error"
// @router /photodetail [get]
func (c *MainController) PhotoDetail() {
	c.Data["Website"] = "beego.vip"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	//c.ViewPath = "test.html"
}
