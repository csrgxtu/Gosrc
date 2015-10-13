package routers

import (
	"github.com/csrgxtu/auth0/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
