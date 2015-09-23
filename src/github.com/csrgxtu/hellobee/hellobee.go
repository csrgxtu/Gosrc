package hellobee

import (
    "github.com/astaxie/beego"
)

type MainController struct {
    beego.Controller
}

func (this *MainController) Get() {
    this.Ctx.WriteString("hello bee")
}

func main() {
    beego.Router("/", &MainController{})
    beego.Run()
}
