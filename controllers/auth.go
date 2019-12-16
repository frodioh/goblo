package controllers

import (
	"net/http"
)

var globalToken = []byte("secret")

func GetToken(w http.ResponseWriter, r *http.Request) {
}