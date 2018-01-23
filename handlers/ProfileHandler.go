package handlers

import (
	"net/http"
	"context"
	"github.com/kartikkh/Medium/controllers"
	"github.com/kartikkh/Medium/models"
	"fmt"
	"github.com/kartikkh/Medium/auth"
)


const (
	CurrentUser    = contextKey("current_user")
	Claim          = contextKey("claim")
)

func (h *Handler) ProfileHandler(w http.ResponseWriter, r *http.Request) {
	user := controllers.Controller(h.DB)
	router := NewRouter()

	router.AddRoute(`profiles\/(?P<username>[0-9a-zA-Z\-]+)$`,"GET", h.getCurrentUser(user.GetProfile))
	router.AddRoute(`profiles\/(?P<username>[0-9a-zA-Z\-]+)\/follow`,"POST", h.getCurrentUser(user.FollowUser))
	router.AddRoute(`profiles\/(?P<username>[0-9a-zA-Z\-]+)\/follow`,"DELETE", h.getCurrentUser(user.UnfollowUser))
	router.ServeHTTP(w, r)

}

func (h *Handler) getCurrentUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var u = &models.User{}
		ctx := r.Context()

		if claim, _ := auth.CheckRequest(r); claim != nil {
			if u, err = h.DB.FindUserByUserName(claim.Username); err != nil {
				http.Error(w, fmt.Sprint("User with username", claim.Username, "doesn't exist !"), http.StatusUnauthorized)
				return
			}
			fmt.Printf("%q\n", claim)
			ctx = context.WithValue(ctx, "claim", claim)
			fmt.Printf("%q\n", ctx)
		}

		ctx = context.WithValue(ctx, "current_user", u)
	//	fmt.Printf("%q\n", "djshfksdhf")
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
}