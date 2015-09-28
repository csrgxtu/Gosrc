package main

import (
    "encoding/json"
    "fmt"
)

type Server struct {
    ServerName string
    ServerIP   string
}

type Serverslice struct {
    Servers []Server
}

type Article struct {
  ArticleID string
  Title string
  Content string
  Author string
  CreatedAt string
}

type Result struct {
  Err bool
  // msg string
  // total int
  Recs []Article
}

func main() {
    var s Serverslice
    s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
    s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
    b1, err := json.Marshal(s)
    if err != nil {
        fmt.Println("json err:", err)
    }
    fmt.Println(string(b1))

    var r Result
    r.Err = true
    // r.msg = "Successful"
    // r.total = 0
    r.Recs = append(r.Recs, Article{ArticleID: "D25FAD", Title: "Think Differently", Content: "Short Content", Author: "Archr", CreatedAt: "2015-09-28 17:20:01"})
    r.Recs = append(r.Recs, Article{ArticleID: "D25FAD", Title: "Think Differently", Content: "Short Content", Author: "Archr", CreatedAt: "2015-09-28 17:20:01"})
    r.Recs = append(r.Recs, Article{ArticleID: "D25FAD", Title: "Think Differently", Content: "Short Content", Author: "Archr", CreatedAt: "2015-09-28 17:20:01"})
    b, err1 := json.Marshal(r)
    if err1 != nil {
      fmt.Println("json err:", err1)
    }
    fmt.Println(string(b))
}
