package main

import (
  "fmt"
  "gopkg.in/mgo.v2"
  "os"
)

func main() {
  var uri = "mongodb://linyy:linyy1234@192.168.200.20/bookshelf"

  sess, err := mgo.Dial(uri)
  if err != nil {
    fmt.Printf("Can't connect to mongo, go error %v\n", err)
    os.Exit(1)
  } else {
    fmt.Printf("established")
  }
  defer sess.Close()
}
