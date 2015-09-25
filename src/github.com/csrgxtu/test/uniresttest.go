package main

import "fmt"
import "github.com/csrgxtu/unirest"

func main() {
  var url = "http://www.douban.com"

  fmt.Println(unirest.DoGet(url))
  // fmt.Println(body)
}
