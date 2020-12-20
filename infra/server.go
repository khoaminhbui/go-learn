package infra

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/khoaminhbui/go-learn/domain"

	"github.com/gorilla/mux"
)

// Response define data structure for http response object
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
		prime := domain.TrialDivision(n)
		response.Message = "Success"
		response.Value = prime
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// StartSimpleServer run a simple http server to handle request
func StartSimpleServer() {
	fmt.Println("Start HTTP server...")

	// Init router
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/prime/{n}", getPrime).Methods("GET")

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Println("Server run at port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
