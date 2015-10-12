package controllers

import (
    "encoding/json"
    "fmt"
    "net/http"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"

    "github.com/julienschmidt/httprouter"
    "github.com/csrgxtu/mgo/models"
)

type (
    // UserController represents the controller for operating on the User resource
    UserController struct {
        session *mgo.Session
    }
)

func NewUserController(s *mgo.Session) *UserController {
    return &UserController{s}
}

// GetUser retrieves an individual user resource
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    // Grab id
    id := p.ByName("id")

    // Verify id is ObjectId, otherwise bail
    if !bson.IsObjectIdHex(id) {
        w.WriteHeader(404)
        return
    }

    // Grab id
    oid := bson.ObjectIdHex(id)

    // Stub user
    u := models.User{}

    // Fetch user
    if err := uc.session.DB("go_rest_tutorial").C("users").FindId(oid).One(&u); err != nil {
        w.WriteHeader(404)
        return
    }

    // Marshal provided interface into JSON structure
    uj, _ := json.Marshal(u)

    // Write content-type, statuscode, payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w, "%s", uj)
}

// CreateUser creates a new user resource
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    // Stub an user to be populated from the body
    u := models.User{}

    // Populate the user data
    json.NewDecoder(r.Body).Decode(&u)

    // Add an Id
    u.Id = bson.NewObjectId()

    // Write the user to mongo
    uc.session.DB("go_rest_tutorial").C("users").Insert(u)

    // Marshal provided interface into JSON structure
    uj, _ := json.Marshal(u)

    // Write content-type, statuscode, payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    fmt.Fprintf(w, "%s", uj)
}

// RemoveUser removes an existing user resource
func (uc UserController) RemoveUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    // TODO: only write status for now
    w.WriteHeader(200)
}
