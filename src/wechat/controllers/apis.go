package controllers

import (
  "github.com/martini-contrib/render"
)

func Create(r render.Render) {
  var rt models.Result

  r.JSON(200, rt)
}
