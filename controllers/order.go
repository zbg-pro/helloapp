package controllers

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"helloapp/models"
)

// Operation about order
type OrderController struct {
	beego.Controller
}

// @Title AddOrder
// @Description add Order object
// @Param	body 	models.Order	true		"The object content"
// @Success 200 {string} models.Order.Id
// @Failure 403 body is empty
// @router /addOrder [post]
func (o *OrderController) AddOrder() {
	var order models.Order

	json.Unmarshal(o.Ctx.Input.RequestBody, &order)
	uid := models.AddOrder(order)
	o.Data["json"] = map[string]string{"uid": uid}
	o.ServeJSON()
}

func (o *OrderController) Test() {
	o.Data["Website"] = "beego.vip"
	o.Data["Email"] = "astaxie@gmail.com"
	o.TplName = "index.tpl"

	/*o.Data["json"] = "logout success"
	o.ServeJSON()*/
}

func (o *OrderController) Test2() {
	o.Ctx.WriteString("aaaaaa")
}
