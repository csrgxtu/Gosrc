package unirest

import "fmt"
import "net/http"
import "io/ioutil"

func DoGet(url string) body string {
  resp, err := http.Get(url)
  if err != nil {
    return nil
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  body = string(body)
  // fmt.Println("%s\n", string(body))
  return body
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
