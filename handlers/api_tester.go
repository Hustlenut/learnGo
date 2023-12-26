package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type IntData struct {
	Number int `json:"number"`
}

//This takes in a number and confirms it in the response
func HandleBasicPost(w http.ResponseWriter, r *http.Request) {
	var data IntData

	decoder := json.NewDecoder(r.Body)
	/*
		Go uses value-copy mechanisme. When passing a variable
		to a function, a copy is made and this wont affect the
		original variable. This "&" pointer allows this modification.

		"nil" is the equivalent to a uninitialized pointer in other languages.

		In Go, you can use short variable declaration:
		if variable := expression; condition {}
	*/
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	// Log the received number
	log.Printf("Received number: %d\n", data.Number)

	// Respond with a success message
	response := map[string]string{"message": "POST request received successfully"}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
