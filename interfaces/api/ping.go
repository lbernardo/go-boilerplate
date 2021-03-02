package api

import (
	"net/http"
	"time"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	t := time.Now().String()
	w.Write([]byte("PING "+ t))
	w.WriteHeader(http.StatusOK)
}
