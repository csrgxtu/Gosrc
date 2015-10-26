package main

import (
  "github.com/go-martini/martini"

  "wechat/controllers"
)

func main() {
  m := martini.Classic()

  m.Get("/apis", controllers.Create)

  m.Run()
}
