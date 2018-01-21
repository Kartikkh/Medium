package handlers

import (
	"net/http"
	"context"
	"github.com/kartikkh/Medium/controllers"
	"github.com/kartikkh/Medium/models"
)


const (
	CurrentUser    = contextKey("current_user")
	FetchedArticle = contextKey("article")
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
		//var err error
		var u = &models.User{}
		ctx := r.Context()

		//if claim, _ := auth.CheckRequest(r); claim != nil {
		//	// Check also that user exists and prevent old token usage
		//	// to gain privillege access.
		//	if u, err = h.DB.FindUserByUsername(claim.Username); err != nil {
		//		http.Error(w, fmt.Sprint("User with username", claim.Username, "doesn't exist !"), http.StatusUnauthorized)
		//		return
		//	}
		//	ctx = context.WithValue(ctx, Claim, claim)
		//}

		ctx = context.WithValue(ctx, CurrentUser, u)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
}