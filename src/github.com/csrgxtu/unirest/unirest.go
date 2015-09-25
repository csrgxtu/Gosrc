package unirest

import "fmt"
import "net/http"
import "io/ioutil"

func DoGet(url string) int {
  resp, err := http.Get(url)
  if err != nil {
    return 500
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  // fmt.Println("%s\n", string(body))
  return 200
}

func DoPut(url string) {
  fmt.Println(url)
}

func DoPost(url string) {
  fmt.Println(url)
}

func DoDelete(url string) {
  fmt.Println(url)
}
