package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type HealthCheck struct {
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp"`
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/health", healthCheckHandler).Methods("GET")

	log.Println("api started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(HealthCheck{Status: "OK",
		Timestamp: time.Now().UnixNano() / 1000000})
}
