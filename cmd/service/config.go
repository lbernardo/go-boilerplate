package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Routes struct {
	route *mux.Router
}


func NewRoutes(r *mux.Router) *Routes {
	return &Routes{
		route: r,
	}
}

type Callback func (w http.ResponseWriter, r *http.Request)

func (r *Routes) Get(path string, callback Callback) {
	r.route.HandleFunc(path, callback).Methods(http.MethodGet)
}

func (r *Routes) Post(path string, callback Callback) {
	r.route.HandleFunc(path, callback).Methods(http.MethodPost)
}

func (r *Routes) Delete(path string, callback Callback) {
	r.route.HandleFunc(path, callback).Methods(http.MethodDelete)
}

func (r *Routes) Group(path string, callback func(r *mux.Router)) {
	router := r.route.PathPrefix(path).Subrouter()
	callback(router)
}