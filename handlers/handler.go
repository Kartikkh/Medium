package handlers

import (
	"gopkg.in/mgo.v2"
	"net/http"
)

type Handler struct {
	session *mgo.Session
}

func NewHandler(s *mgo.Session) *Handler {
	return &Handler{s}
}


func (h *Handler) UsersHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":

	case "GET":
		// TODO:
		// Check auth
		// Get current users
	case "PUT":
		// TODO:
		// Check auth
		// Update user
	default:
		http.NotFound(w, r)
	}
}



func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
	//	h.LoginUser(w, r)
	default:
		http.NotFound(w, r)
	}
}
