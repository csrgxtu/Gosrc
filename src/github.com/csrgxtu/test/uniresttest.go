package main

import "github.com/csrgxtu/unirest"

func main() {
  var url = "http://www.douban.com"

  unirest.DoGet(url)
}
