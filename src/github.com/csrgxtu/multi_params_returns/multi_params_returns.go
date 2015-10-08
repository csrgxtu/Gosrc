package main

import "fmt"

func swap(a int, b int) (reta int, retb int) {
  reta = b
  retb = a
  return
}

func main() {
  a, b := swap(3, 4)
  fmt.Println(a)
  fmt.Println(b)
}
