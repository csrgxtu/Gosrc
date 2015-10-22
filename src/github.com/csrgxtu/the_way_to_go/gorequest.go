package main

import "fmt"
import "github.com/parnurzeal/gorequest"

func main() {
	request := gorequest.New()
	resp, body, errs := request.Get("http://www.douban.com").End()
	if errs != nil {
		fmt.Println(errs)
	}
	fmt.Println(resp)
	fmt.Println(body)
}
