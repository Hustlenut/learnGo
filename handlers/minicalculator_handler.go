package handlers

import (
	"encoding/json"
	"net/http"
	"sync"

	"learnGo/models"
)

var (
	/*
		A mutex (short for "mutual exclusion") is a
		synchronization primitive used in concurrent
		programming to protect shared data from being
		accessed simultaneously by multiple goroutines
		(concurrent functions or threads).
		Mutexes ensure that only one goroutine can access
		a critical section of code at a time, preventing
		data races and maintaining data consistency.
	*/
	mu      sync.Mutex
	result  models.CalculationResult
	hasData bool
)

// PostNumbers handles the POST request to perform a mathematical operation.
func PostNumbers(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var numbers models.Number
	err := decoder.Decode(&numbers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Perform the calculation with numbers.Value
	result = models.CalculationResult{
		Operation: "addition",
		Result:    numbers.Value + numbers.Value,
	}

	// Set the hasData flag to true
	mu.Lock()
	hasData = true
	mu.Unlock()

	// Respond with the result
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func GetAnswer(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock() //used to schedule a function call to be executed just before the surrounding function returns

	if !hasData {
		http.Error(w, "No result available", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
