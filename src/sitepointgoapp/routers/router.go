package routers

import (
	"sitepointgoapp/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
		beego.Router("/login", &controllers.MainController{}, "get:HelloSitepoint")
}
