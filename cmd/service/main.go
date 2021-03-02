package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()
	NewRoutes(r).Register()
	log.Println(http.ListenAndServe(fmt.Sprintf("%v:%v", os.Getenv("HOST"), os.Getenv("PORT")), r))
}
