package controllers

import (
	"net/http"
	"encoding/json"
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
	defer r.Body.Close()
	u := &body.User
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		logrus.Info(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	m ,err:= models.NewUser(u.Email,u.Username,u.Password)
	if err != nil {
		logrus.Info(err)
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = uc.DB.CreateUser(m)

	if err != nil {
		logrus.Error(err)
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

	body := struct {
		User struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		} `json:"user"`
	}{}

	defer r.Body.Close()
	u := &body.User
	err := json.NewDecoder(r.Body).Decode(&body)


	if err != nil {
		logrus.Fatal(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	m, err := uc.DB.FindUserByEmail(u.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	print("3")
	match := m.MatchPassword(u.Password)
	if !match {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201

	res := &UserJSON{
		&User{
			Username: m.Username,
			Email:    m.Email,
			Token:   auth.GetToken(m.Username),
			Bio:      m.Bio,
			Image:    m.Image,
		},
	}
	json.NewEncoder(w).Encode(res)
}



func (uc *UserController) GetCurrentUser(w http.ResponseWriter, r *http.Request) {

	claim , err :=  auth.CheckRequest(r)
	if err!=nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	m, err := uc.DB.FindUserByUserName(claim.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	res := &UserJSON{
		&User{
			Username: m.Username,
			Email:    m.Email,
			Token:    auth.GetToken(m.Username),
			Bio:      m.Bio,
			Image:    m.Image,
		},
	}
	json.NewEncoder(w).Encode(res)


}

func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {

	claim , err :=  auth.CheckRequest(r)

	 if err!=nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	 }
	body := struct {
		User struct {
			Email    string `json:"email"`
			Password string `json:"password"`
			Image string    `json:"image"`
			Bio string      `json:"bio"`
			Username string `json:"username"`
		} `json:"user"`
	}{}

	defer r.Body.Close()
	u := &body.User
	error := json.NewDecoder(r.Body).Decode(&body)

	if error != nil {
		logrus.Info(error)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}


	m , err := uc.DB.FindUserByUserName(claim.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}


	if u.Username != ""{
		m.Username = u.Username
	}
	if u.Password != ""{
		m.Password = models.EncryptPassword(u.Password)
	}
	if u.Image != ""{
		m.Image = u.Image
	}
	if u.Bio != ""{
		m.Bio = u.Bio
	}
	if u.Email != ""{
		m.Email = u.Email
	}
	uc.DB.Save(&m)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)


	res := &UserJSON{
		&User{
			Username: m.Username,
			Email:    m.Email,
			Token:   auth.GetToken(m.Username),
			Bio:      m.Bio,
			Image:    m.Image,
		},
	}
	json.NewEncoder(w).Encode(res)



}