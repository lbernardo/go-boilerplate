package main

import (
	"github.com/lbernardo/go-boilerplate/interfaces/api"
)

func (r *Routes) Register() {
	r.Get("/ping", api.Ping)
}
