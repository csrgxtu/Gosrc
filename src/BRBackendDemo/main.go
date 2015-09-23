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
	this.Data["json"] = "{fuck}"
	this.ServeJson()
  // this.Ctx.WriteString("yo, man")
}

func main() {
	beego.Router("/article", &MainController{})
	beego.Run()
}
