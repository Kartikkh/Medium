package controllers

import (
	"net/http"
	"encoding/json"
	"fmt"

	"github.com/kartikkh/Medium/auth"
	"github.com/kartikkh/Medium/models"
)

type Profile struct {
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
	Following bool 	`json:"following"`
}

type ProfileJSON struct {
	Profile *Profile `json:"profile"`
}


func (uc *UserController) GetProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
    username := ctx.Value("username").(string)
	LoginUser := ctx.Value("current_user").(*models.User)

    m , err := uc.DB.FindUserByUserName(username)
    if err != nil{
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	fmt.Printf("%q\n", m)
	res := &ProfileJSON{
		&Profile{
			Username: m.Username,
			Bio:      m.Bio,
			Image:    m.Image,
			Following: uc.DB.IsFollowing(m, LoginUser),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	json.NewEncoder(w).Encode(res)
}


func(uc *UserController) FollowUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	res := "followUser"
	json.NewEncoder(w).Encode(res)
}


func (uc *UserController) UnfollowUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	res := "UnfollowUser"
	json.NewEncoder(w).Encode(res)

}

