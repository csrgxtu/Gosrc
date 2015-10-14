package main

import "fmt"

func main() {
  mf := map[int]func() int{
    1: func() int {return 10},
    2: func() int {return 20},
    3: func() float32 {return 30.0},
    5: func() int {return 50},
  }
  fmt.Println(mf)
}
