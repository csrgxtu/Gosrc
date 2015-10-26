package main

import (
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"

  "wechat/controllers"
)

func main() {
  m := martini.Classic()

  m.Use(render.Renderer())

  m.Get("/apis", controllers.Create)

  m.Run()
}
