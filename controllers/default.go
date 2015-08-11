package controllers

import (
	"Go/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["IsHome"] = true
	c.TplNames = "home.html"
	var err error
	c.Data["blogs"], err = models.GetAllTopics()
	if err != nil {
		beego.Error(err)
	}
	c.Data["IsLogin"] = checkAccount(c.Ctx)
}
