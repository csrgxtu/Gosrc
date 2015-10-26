package main

import (
	"github.com/go-martini/martini"
	// "github.com/martini-contrib/oauth2"
	"github.com/csrgxtu/oauth2"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	goauth2 "golang.org/x/oauth2"

	"oauth/controllers"
)

func main() {
	m := martini.Classic()

	m.Use(render.Renderer())
	m.Use(sessions.Sessions("my_session", sessions.NewCookieStore([]byte("secret123"))))
	// m.Use(oauth2.Google(
	//   &goauth2.Config{
	//     ClientID:     "393837115699-selpvjkb9bl6orqar3915dpb142u5u27.apps.googleusercontent.com",
	//     ClientSecret: "CQm04gHSRATBsZrXVqVXmcj6",
	//     Scopes:       []string{"https://www.googleapis.com/auth/drive"},
	//     RedirectURL:  "http://localhost:3000/apis",
	//   },
	// ))
	// m.Use(oauth2.Github(
	// 	&goauth2.Config{
	// 		ClientID:     "855acacb3bfe44dd929b",
	// 		ClientSecret: "da676b7e863cf493f9fcac0eb856cbb2074c5806",
	// 		RedirectURL:  "http://localhost:3000/apis",
	// 	},
	// ))
	m.Use(oauth2.WeChat(
		&goauth2.Config{
			ClientID: "wxa3b6a8fc961fef35",
			ClientSecret: "6790ebc1a1cd60b5f5d2c5bfc0bf4391",
			Scopes: []string{"snsapi_login"},
			RedirectURL: "http://dev.beautifulreading.com",
		},
	))

	m.Get("/apis", controllers.Create)
	m.Get("/", controllers.Welcome)
	m.Get("/auth", oauth2.LoginRequired, func(tokens oauth2.Tokens) string {
		return tokens.Access()
	})

	// m.Run()
	m.RunOnAddr(":3002")
}
