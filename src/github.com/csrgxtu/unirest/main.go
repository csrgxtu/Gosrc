package main

// import "fmt"
// import "net/http"
// import "os"
// import "io/ioutil"
import "./unirest"

// func doGet(url string) {
//   fmt.Println(url)
// }

func main() {
  // resp, err := http.Get("http://book.douban.com/")
  // if err != nil {
  //   fmt.Println("Http Get Error")
  //   os.Exit(1)
  // }
  // defer resp.Body.Close()
  // // fmt.Println(&resp.Body)
  // body, err := ioutil.ReadAll(resp.Body)
  // if err != nil {
  //   fmt.Println("ioutil read error")
  //   os.Exit(1)
  // }
  // fmt.Println("%s\n", string(body))
  doGet("http://www.douban.com")
}
