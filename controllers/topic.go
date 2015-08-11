package controllers

import (
	"Go/models"
	// "fmt"
	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsTopic"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	var err error
	this.Data["Topics"], err = models.GetAllTopics()
	if err != nil {
		beego.Error(err)
	}
	this.TplNames = "topic.html"
}

func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	tid := this.Input().Get("tid")
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	var err error
	if len(tid) != 0 {
		err = models.ModifyTopic(tid, title, content)
	} else {
		err = models.AddTopic(title, content)
	}
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic", 302)
}

func (this *TopicController) Add() {
	// this.Ctx.WriteString("add")
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.TplNames = "topic_add.html"
}
func (this *TopicController) Modify() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	if this.Data["IsLogin"] == false {
		this.Redirect("/login", 302)
		return
	}
	var err error
	tid := this.Input().Get("tid")
	t, err := models.GetTopic(tid)
	this.Data["Blog"] = t
	if err != nil {
		beego.Error(err)
	}
	this.TplNames = "topic_modify.html"
}
