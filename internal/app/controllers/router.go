package controllers

import "net/http"

//Router ..
type Router struct {
	router *http.ServeMux
}

//NewRouter ..
func NewRouter() *Router {
	return &Router{
		router: http.NewServeMux(),
	}
}
