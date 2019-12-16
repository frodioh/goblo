package main

import (
	"github.com/frodion/goblo/controllers"
	"github.com/frodion/goblo/models"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func main() {
	models.InitDB("user=radiorodeo password=gjfdksg2302j dbname=radiorodeo sslmode=disable")

	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./views/")))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

//	r.Handle("/get-token", controllers.GetToken).Methods("GET")
	r.HandleFunc("/status", controllers.GetStatus).Methods("GET")
	r.HandleFunc("/posts", controllers.GetPosts).Methods("GET")

	err := http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))

	if err == nil {
		println(err)
	}
}