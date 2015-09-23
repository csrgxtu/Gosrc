package main

import (
	"github.com/astaxie/beego"
	// "encoding/json"
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
	// mapB, _ := json.Marshal(mapD)
	this.Data["json"] = mapD
	this.ServeJson()
  // this.Ctx.WriteString("yo, man")
}

func main() {
	beego.Router("/article", &MainController{})
	beego.Run()
}
