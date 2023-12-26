package config

import (
	"net/http"

	"learnGo/handlers"

	"github.com/gorilla/mux"
)

func ConfigBasicPost() {
	r := mux.NewRouter()

	r.HandleFunc("/basicapi", handlers.HandleBasicPost).Methods("POST")

	http.ListenAndServe(":9999", r)
}

func ConfigMinicalcRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/minicalc", handlers.PostNumbers).Methods("POST")
	r.HandleFunc("/minicalc", handlers.GetAnswer).Methods("GET")

	http.ListenAndServe(":8080", r)
}
