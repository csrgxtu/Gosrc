package main

import (
	_ "github.com/csrgxtu/bapi/docs"
	_ "github.com/csrgxtu/bapi/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
