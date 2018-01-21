package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/sirupsen/logrus"
)

func (uc *UserController) GetProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	logrus.Info(r.RequestURI)
	res := "GetProfile"
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

