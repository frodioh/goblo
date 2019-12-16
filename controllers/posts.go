package controllers

import (
	"encoding/json"
	"github.com/frodion/goblo/models"
	"net/http"
)

func GetPosts(w http.ResponseWriter, r *http.Request){
	payload, _ := json.Marshal(models.GetPosts())

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
}