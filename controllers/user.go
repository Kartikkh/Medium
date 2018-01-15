package controllers

import (
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/kartikkh/Medium/models"
	"github.com/sirupsen/logrus"
	"github.com/kartikkh/Medium/auth"
)


type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}

type UserJSON struct {
	User *User `json:"user"`
}


type UserController struct {
	DB    *models.DB
}


func Controller(db *models.DB) *UserController {
	return &UserController{db}
}


func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {

	body := struct {
		User struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Password string `json:"password"`
		} `json:"user"`
	}{}

	u := &body.User
	err := json.NewDecoder(r.Body).Decode(&body)


	if err != nil {
		logrus.Info(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Println(u)

	m ,err:= models.NewUser(u.Email,u.Username,u.Password)
	if err != nil {
		logrus.Info(err)
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = uc.DB.CreateUser(m)
	if err != nil {
		logrus.Error(err)
		// TODO: Error JSON
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	res := &UserJSON{
		&User{
			Username: m.Username,
			Email:    m.Email,
			Token : auth.GetToken(m.Username),
		},
	}
	json.NewEncoder(w).Encode(res)

}



func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {



}



func (uc *UserController) GetCurrentUser(w http.ResponseWriter, r *http.Request) {



}

func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {



}