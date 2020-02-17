package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type HealthCheck struct {
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp"`
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/health", HealthCheckHandler).Methods("GET")

	log.Println("api started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(HealthCheck{Status: "OK",
		Timestamp: time.Now().UnixNano() / 1000000})
}
