package main

import (
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  "github.com/martini-contrib/sessions"
  "github.com/martini-contrib/oauth2"
  goauth2 "golang.org/x/oauth2"

  "oauth/controllers"
)

func main() {
  m := martini.Classic()

  m.Use(render.Renderer())
  m.Use(sessions.Sessions("my_session", sessions.NewCookieStore([]byte("secret123"))))
  m.Use(oauth2.Google(
    &goauth2.Config{
      ClientID:     "393837115699-selpvjkb9bl6orqar3915dpb142u5u27.apps.googleusercontent.com",
      ClientSecret: "CQm04gHSRATBsZrXVqVXmcj6",
      Scopes:       []string{"https://www.googleapis.com/auth/drive"},
      RedirectURL:  "http://localhost:3000/apis",
    },
  ))

  m.Get("/apis", controllers.Create)
  m.Get("/", controllers.Welcome)
  m.Get("/auth", oauth2.LoginRequired, func(tokens oauth2.Tokens) string {
    return tokens.Access()
  })

  m.Run()
}
