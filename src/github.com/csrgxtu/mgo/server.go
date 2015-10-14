package main

import (
    // Standard library packages
    "net/http"

    // Third party packages
    "github.com/julienschmidt/httprouter"
    "github.com/csrgxtu/mgo/controllers"
    "gopkg.in/mgo.v2"

    "fmt"
)

func main() {
    // Instantiate a new router
    r := httprouter.New()

    // Get a UserController instance
    uc := controllers.NewUserController(getSession())

    // Get a user resource
    r.GET("/user/:id", uc.GetUser)

    r.POST("/user", uc.CreateUser)

    r.DELETE("/user/:id", uc.RemoveUser)

    // Fire up the server
    fmt.Println("Server started at localhost:3000")
    http.ListenAndServe("localhost:3000", r)
}

func getSession() *mgo.Session {
    // Connect to our local mongo
    s, err := mgo.Dial("mongodb://localhost")

    // Check if connection error, is mongo running?
    if err != nil {
        panic(err)
    }
    return s
}
