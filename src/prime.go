package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Response struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
	Value     int    `json:"value"`
}

func trialDivision(n int) int {
	for i := n - 1; i >= 2; i-- {
		var isPrime = true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			return i
		}
	}

	return 1
}

func getPrime(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	n, err := strconv.Atoi(params["n"])

	response := Response{ErrorCode: 0, Message: "", Value: 0}
	if err != nil {
		response.ErrorCode = 1
		response.Message = "Invalid number"
	} else {
		prime := trialDivision(n)
		response.Message = "Success"
		response.Value = prime
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	fmt.Println("hello prime!")

	// Init router
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/prime/{n}", getPrime).Methods("GET")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
