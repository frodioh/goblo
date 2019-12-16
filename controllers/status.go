package controllers

import (
	"net/http"
)

func GetStatus(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
}