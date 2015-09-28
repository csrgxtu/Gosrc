package routers

import (
	"github.com/csrgxtu/mysqlTut/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
