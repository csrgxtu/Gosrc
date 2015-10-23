package main

import (
  jwt "github.com/dgrijalva/jwt-go"
  "io/ioutil"
  "net/http"
  "fmt"
  "time"
  "log"
  "io"
  "os"
  // "strings"
)

var (
  privateKey []byte
  publicKey []byte
)

func init() {
  privateKey, _ = ioutil.ReadFile("demo.rsa")
  publicKey, _ = ioutil.ReadFile("demo.rsa.pub")
}

func main() {
  http.HandleFunc("/", handleRequest)
  http.HandleFunc("/auth", authRequest)
  http.HandleFunc("/users", GetUserInfo)

  fmt.Print("JWT Server Listen On Port 3000 ...\n")

  err := http.ListenAndServe(":3000", nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  } else {
    // log.Info("JWT Server Listen On Port 3000 ...")
    // fmt.Print("JWT Server Listen On Port 3000 ...")
  }
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
  token := jwt.New(jwt.GetSigningMethod("RS256")) // Create a Token that will be signed with RSA 256.
  /*
  {
      "typ":"JWT",
      "alg":"RS256"
  }
  */

  token.Claims["ID"] = "This is my super fake ID"
  token.Claims["exp"] = time.Now().Unix() + 36000

  // The claims object allows you to store information in the actual token.

  tokenString, _ := token.SignedString(privateKey)

  // tokenString Contains the actual token you should share with your client.

  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, tokenString)
}

func GetUserInfo(res http.ResponseWriter, req *http.Request) {
  // AuthKey := http.CanonicalHeaderKey("authorization")
  // tokenString := strings.TrimLeft(req.Header.Get(AuthKey), "Bearer ")
  tokenString := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6IlRoaXMgaXMgbXkgc3VwZXIgZmFrZSBJRCIsImV4cCI6MTQ0NTYxMzM5N30.BY4IDUPrkyAu321fdeE_bsGvyee_my0HcRub2hrGMxUqWBshc_5jrRSHXN3bdX6Qx8ThB2BRM5jwaOPvaFaO4ThrF7R7BjxnzvlPzlo8NcVeiQWe07E992gywnI6dU-BExkKFSE0zZSb37w0D0bJn0-575KzqgeV4MfKwdR8aSw"

  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    fmt.Println(token)
    // return myLookupKey(token.Header["kid"])
    // return nil, nil
    return loadData("demo.rsa.pub")
  })

  if err == nil && token.Valid {
      fmt.Println("Your token is valid.  I like your style.")
  } else {
      fmt.Println("This token is terrible!  I cannot accept this.")
      fmt.Println(err)
  }

  res.WriteHeader(http.StatusOK)
  fmt.Fprint(res, tokenString)
}

func authRequest(w http.ResponseWriter, r *http.Request) {
  token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (face interface{}, err error) {
    fmt.Println(face)
    return face, nil
    // return publicKey, nil
  })
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(token)

  // if token.Valid {
  //   //YAY!
  //   fmt.Println("token Valid")
  // } else {
  //   //Someone is being funny
  //   fmt.Println("totken Unautherized")
  // }
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
