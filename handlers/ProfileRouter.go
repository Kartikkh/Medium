package handlers

import (
	"net/http"

)

type contextKey string

type Route struct {
	Pattern string
	ActionHandlers map[string]http.Handler
}


type Router struct {
	http.Handler
	routes []Route
}


func NewRouter() *Router {
	return &Router{
		routes: make([]Route, 0),
	}
}

func (r *Router) AddRoute(pattern string, method string, handler http.Handler) {

}

