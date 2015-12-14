package main

import "time"
import "fmt"

func main() {
	bla := time.Now()
	fmt.Printf("%d\n", bla.Unix())
}
