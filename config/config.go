package config

import (
	"log"
	"net/http"

	"learnGo/handlers"

	"github.com/gorilla/mux"
)

func HandleMinicalcRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/minicalc", handlers.PostNumbers).Methods("POST")
	r.HandleFunc("/minicalc", handlers.GetAnswer).Methods("GET")

	log.Fatal(http.ListenAndServe(":9000", r))
}
