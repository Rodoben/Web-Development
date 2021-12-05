package controllers

import (
	"controllers/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc UserController) Getuser(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "Ronald Benjamin",
		Gender: "male",
		Age:    32,
		Id:     p.ByName("id"),
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(res, "%s\n", uj)
}

func (uc UserController) CreateUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	u := models.User{}
	json.NewDecoder(req.Body).Decode(&u)
	fmt.Println("decode done")
	u.Id = "007"
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("marshal done")
	//res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(res, "%s\n", uj)

}
func (uc UserController) DeleteUser(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	res.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(res, "Write code to delete user\n")
}
