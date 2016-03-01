package main

import (
	"fmt"
	"time"
)

func main() {
	//var layout = "1970-01-01 15:04"
	//var str = "2016-3-01"
	//t, err := time.Parse("2006-01-02 15:04", "2011-01-19 22:15")
	t, err := time.Parse("2006-01-02", "2016-03-01")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(t.Month())
}
