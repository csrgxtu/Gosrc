package main

import (
	"github.com/astaxie/beego"
)

type ResData struct {
	code int
	msg string
}

type MainController struct {
  beego.Controller
}

func (this *MainController) Get() {
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	this.Data["json"] = mapD
	this.ServeJson()
}

func main() {
	beego.Router("/article", &MainController{})
	beego.Run()
}
