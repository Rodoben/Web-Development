package controllers

import (
	"encoding/json"
	"ex2/models"
	"fmt"
	"net/http"

	uuid "satori/go.uuid"

	"github.com/julienschmidt/httprouter"
)

type UserController struct {
	session map[string]models.User
}

func NewUserController(m map[string]models.User) *UserController {
	return &UserController{m}
}

func (uc UserController) GetUser(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	u := uc.session[id]
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(res, "%s\n", uj)

}

func (uc UserController) CreateUser(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := models.User{}
	json.NewDecoder(req.Body).Decode(&u)
	u.Id = uuid.NewV4().String()
	uc.session[u.Id] = u
	models.StoreUsers(uc.session)
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(res, "%s\n", uj)

}

func (uc UserController) DeleteUser(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	delete(uc.session, id)
	models.StoreUsers(uc.session)
	res.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(res, "Deleted user ", id, "\n")
}
