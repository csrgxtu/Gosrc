package main

import (
  "fmt"
  "time"
)

var week time.Duration

func main() {
  t := time.Now()
  fmt.Println(t)  // e.g. Wed Dec 21 09:52:14 +0100 RST 2011
  fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
  // 21.12.2011
  t = time.Now().UTC()
  fmt.Println(t) // Wed Dec 21 08:52:14 +0000 UTC 2011
  fmt.Println(time.Now()) // Wed Dec 21 09:52:14 +0100 RST 2011
}
