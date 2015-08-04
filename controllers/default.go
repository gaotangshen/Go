package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "index.tpl"
	c.Data["TrueCond"] = true
	c.Data["FalseCond"] = false

	type u struct {
		Name   string
		Age    int
		Gender string
	}
	user := u{
		Name:   "Gaotang",
		Age:    10,
		Gender: "male",
	}
	c.Data["User"] = user
	num := []int{1, 2, 3, 4, 5}
	c.Data["Number"] = num

}
