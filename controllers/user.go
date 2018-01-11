package controllers

import (
	"gopkg.in/mgo.v2"
	"net/http"
)

type UserController struct {
	session *mgo.Session
}

func Controller(s *mgo.Session) *UserController {
	return &UserController{s}
}




func (u *UserController) Register(w http.ResponseWriter, r *http.Request) {





}




func (u *UserController) Login(w http.ResponseWriter, r *http.Request) {



}