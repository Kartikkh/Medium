package handlers

import (
	"net/http"
	"context"
	"regexp"

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

	found :=false
	for _ , route := range r.routes{
		 if route.Pattern == pattern{
		 	found = true
		 	route.ActionHandlers[method] = handler
		 }
	}

	if !found {
		r.routes = append(r.routes, Route{
			Pattern: pattern,
			ActionHandlers: map[string]http.Handler{
				method: handler,
			},
		})
	}

}


func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range router.routes {
		if matched, _ := regexp.MatchString(route.Pattern, r.URL.Path); matched {
			if h, registered := route.ActionHandlers[r.Method]; registered {

				r = r.WithContext(buildContext(route.Pattern, r))
				h.ServeHTTP(w, r)
			} else {
				http.NotFound(w, r)
			}
			return
		}
	}
}


func buildContext(pattern string, r *http.Request) context.Context {
	re := regexp.MustCompile(pattern) // "profiles\\/(?P<username>[0-9a-zA-Z\\-]+)$"

	n1 := re.SubexpNames()  //["","username"]
	//fmt.Printf("%q\n", re.SubexpNames())
	r2 := re.FindAllStringSubmatch(r.URL.Path, -1) //[["profiles/kartik" "kartik"]]
	//fmt.Printf("%q\n", r2)
	//fmt.Printf("%d\n", len(r2))

	ctx := r.Context()

	if len(r2) > 0 {
		for i, n := range r2[0] {
			if n1[i] != "" {
				ctx = context.WithValue(ctx, n1[i], n)
			}
		}
	}
	return ctx
}
