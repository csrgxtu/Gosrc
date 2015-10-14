package main

import "fmt"

func main() {
  src := [10]byte{'h', 'e', 'l', 'l', 'o'}
  fmt.Println(len(src), src)

  dst := string(src[:])
  fmt.Println(len(dst), dst)
}
