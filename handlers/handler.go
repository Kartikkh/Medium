package handlers


import (
	"gopkg.in/mgo.v2"
	"net/http"
	"github.com/kartikkh/Medium/controllers"
)

type Handler struct {
	session *mgo.Session
}

func NewHandler(s *mgo.Session) *Handler {
	return &Handler{s}
}


func (h *Handler) UsersHandler(w http.ResponseWriter, r *http.Request) {
	user := controllers.Controller(h.session)

	switch r.Method {
	case "POST":

	case "GET":
		// TODO:
		// Check auth
		// Get current users
		user.GetCurrentUser(w,r)
	case "PUT":
		// TODO:
		// Check auth
		// Update user
		user.UpdateUser(w,r)

	default:
		http.NotFound(w, r)
	}
}



func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	user := controllers.Controller(h.session)
	switch r.Method {
	case "POST":
		user.Login(w,r)
	default:
		http.NotFound(w, r)
	}
}





func (h *Handler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	user := controllers.Controller(h.session)
	switch r.Method {
	case "POST":
		user.Register(w, r)
	default:
		http.NotFound(w, r)
	}
}