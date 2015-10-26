package controllers

import (
  "github.com/martini-contrib/render"

  "oauth/models"
)

func Create(r render.Render) {
  var rt models.Result

  r.JSON(200, rt)
}

func Welcome(r render.Render) {
  var rt models.Result

  r.JSON(200, rt)
}
