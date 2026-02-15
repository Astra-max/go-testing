package handlers

import (
	"fmt"
	"net/http"
	"variables"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world"))
}