package infra

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/khoaminhbui/go-learn/libs"

	"github.com/gorilla/mux"
)

type Response struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
	Value     int    `json:"value"`
}

func getPrime(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	n, err := strconv.Atoi(params["n"])

	response := Response{ErrorCode: 0, Message: "", Value: 0}
	if err != nil {
		response.ErrorCode = 1
		response.Message = "Invalid number"
	} else {
		prime := libs.TrialDivision(n)
		response.Message = "Success"
		response.Value = prime
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// StartSimpleServer run a simple http server to handle request
func StartSimpleServer() {
	fmt.Print("Start HTTP server...")

	// Init router
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/prime/{n}", getPrime).Methods("GET")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
