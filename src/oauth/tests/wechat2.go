package main

import (
  "fmt"

  "github.com/chanxuehong/wechat/open/oauth2"
  "net/url"
)

var appid = ""
var appsecret = ""
var redirectURL = "dev.beautifulreading.com"
var scopes = "snsapi_userinfo"
var oauth2Config = oauth2.NewConfig(appid, appsecret, redirectURL, scopes)

func main() {
  var state = "what"
  v := url.Values{}
  v.Set("name", "Ava")

  str := oauth2.AuthCodeURL(appid, redirectURL, scopes, state, v)
  fmt.Println(str)
}
