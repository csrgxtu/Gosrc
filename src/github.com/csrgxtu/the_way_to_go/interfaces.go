package main

import "fmt"

type Shaper interface {
  Area() float32
}

type Square struct {
  side float32
}

func (sq *Square) Area() float32 {
  return sq.side * sq.side
}

func main() {
  sql1 := new(Square)
  sql1.side = 5

  //areaIntf := sql1
  fmt.Printf("The square has area: %f\n", sql1.Area())
}
