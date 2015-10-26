package controllers

import (
  "github.com/martini-contrib/render"

  "wechat/models"
)

func Create(r render.Render) {
  var rt models.Result

  r.JSON(200, rt)
}

func Welcome(r render.Render) {
  var rt models.Result

  r.JSON(200, rt)
}
