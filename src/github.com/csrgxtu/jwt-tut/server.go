package main

import (
  jwt "github.com/dgrijalva/jwt-go"
  "io/ioutil"
  "net/http"
  "encoding/json"
  "io"
  "os"
  "fmt"
  "time"
  "log"
  "strings"
)

var (
  privateKey []byte
  publicKey []byte
)

type Result struct {
  Code int `json:"code"`
  Msg string `json:"msg"`
  Data string `json:"data"`
}

func main() {
  http.HandleFunc("/auth", Auth)
  http.HandleFunc("/users", GetUserInfo)

  fmt.Print("JWT Server Listen On Port 3000 ...\n")
  err := http.ListenAndServe(":3000", nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}

func Auth(res http.ResponseWriter, req *http.Request) {
  token := jwt.New(jwt.GetSigningMethod("RS256"))

  token.Claims["ID"] = "This is my super fake ID"
  token.Claims["exp"] = time.Now().Unix() + 36000

  tokenString, _ := token.SignedString(privateKey)

  rt, _ := json.Marshal(Result{http.StatusOK, "Successful", tokenString})
  res.Header().Set("Content-Type", "application/json")
  res.WriteHeader(http.StatusOK)
  res.Write(rt)
}

func GetUserInfo(res http.ResponseWriter, req *http.Request) {
  AuthKey := http.CanonicalHeaderKey("authorization")
  tokenString := strings.Replace(req.Header.Get(AuthKey), "Bearer ", "", -1)
  // fmt.Println(tokenString)

  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    fmt.Println(token)
    return loadData("demo.rsa.pub")
  })

  if err == nil && token.Valid {
      fmt.Println("Your token is valid.  I like your style.")

      rt, _ := json.Marshal(Result{http.StatusOK, "Successful", tokenString})
      res.Header().Set("Content-Type", "application/json")
      res.WriteHeader(http.StatusOK)
      res.Write(rt)
  } else {
      fmt.Println("This token is terrible!  I cannot accept this.")
      fmt.Println(err)

      rt, _ := json.Marshal(Result{http.StatusUnauthorized, "Token Unauthorized", tokenString})
      res.Header().Set("Content-Type", "application/json")
      res.WriteHeader(http.StatusUnauthorized)
      res.Write(rt)
  }
}

func init() {
  privateKey, _ = ioutil.ReadFile("demo.rsa")
  publicKey, _ = ioutil.ReadFile("demo.rsa.pub")
}

// Helper func:  Read input from specified file or stdin
func loadData(p string) ([]byte, error) {
	if p == "" {
		return nil, fmt.Errorf("No path specified")
	}

	var rdr io.Reader
	if p == "-" {
		rdr = os.Stdin
	} else {
		if f, err := os.Open(p); err == nil {
			rdr = f
			defer f.Close()
		} else {
			return nil, err
		}
	}
	return ioutil.ReadAll(rdr)
}
