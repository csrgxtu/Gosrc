package main

import "github.com/parnurzeal/gorequest"
import "fmt"

func main() {
	request := gorequest.New()
	_, body, errs := request.Post("https://api-rio.beautifulreading.com/beautilfulreading/bs/article").
		Set("Authorization", "Basic bG9zZXI6ZW5nbGFuZA==").
		Set("Content-Type", "application/json").
		Send(`{"number": 121, "title":"test", "author": "archer", "content": "what the"}`).
		End()
	if len(errs) != 0 {
		fmt.Printf("errors")
	}
	fmt.Printf(body)
	//fmt.Printf(resp)
}
