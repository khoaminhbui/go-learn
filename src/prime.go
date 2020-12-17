package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
)

func getPrime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(2);
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