package main

import (
	_ "BRBackendDemo/routers"
	"github.com/astaxie/beego"
)

// type Article struct {
// 		id string
// 		title string
// 		content string
// 		author string
// 		created_at string
// }

// type ResData struct {
// 	err int
// 	msg string
// 	// data Article
// }

type MainController struct {
  beego.Controller
}

func (this *MainController) Get() {
	// article := Article{
	// 	id: "d23fsd87sfd",
	// 	title: "Think Differently",
	// 	content: "In this book mainly told u that the truth about the everythings...",
	// 	author: "Archer Reilly",
	// 	created_at: "2015-09-23 20:00:00",
	// }
	// mapD := &ResData{
	// 	err: 0,
	// 	msg: "Successful"}
	// 	// data: []Article{article, article, article}}
	// mapD := map[string]int{"apple": 5, "lettuce": 7}
	// mapD := map[string][]string{"data": ["archer", "dishy"]}
	mapD := map[string]string{"id": "3df89ghsd", "title": "Think Differently", "content": "In this book mainly told u that the truth about the everythings...", "author": "Archer", "created_at": "2015-09-23 21:00:00"}
	this.Data["json"] = mapD
	this.ServeJson()
}

func main() {
	beego.Router("/", &MainController{})
	beego.Router("/article", &MainController{})
	beego.Run()
}
