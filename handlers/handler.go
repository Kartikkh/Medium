package handlers


import (
	"net/http"
	"github.com/kartikkh/Medium/controllers"
	"github.com/kartikkh/Medium/models"
)

type Handler struct {
	DB     *models.DB
}

func NewHandler(db *models.DB) *Handler {
	return &Handler{db}
}


func (h *Handler) UsersHandler(w http.ResponseWriter, r *http.Request) {

	user := controllers.Controller(h.DB)

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
	user := controllers.Controller(h.DB)
	switch r.Method {
	case "POST":
		user.Login(w,r)
	default:
		http.NotFound(w, r)
	}
}





func (h *Handler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	user := controllers.Controller(h.DB)
	switch r.Method {
	case "POST":
		user.Register(w, r)
	default:
		http.NotFound(w, r)
	}
}