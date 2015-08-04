package main

import (
	// "io"
	// "log"
	// "net/http"
	"github.com/astaxie/beego"
	_ "hello/routers"
)

// type MainController struct {
// 	beego.Controller
// }

func main() {
	// Router
	// beego.Router("/", &MainController{})
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
