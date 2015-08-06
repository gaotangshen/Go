package main

import (
	// "io"
	// "log"
	// "net/http"
	// _ "Go/routers"
	"Go/controllers"
	"Go/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()
}

// type MainController struct {
// 	beego.Controller
// }

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	// Router
	beego.Router("/", &controllers.MainController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Run()
	// http.HandleFunc("/", sayHello)
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

// func sayHello(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, "Hello. world")
// }
