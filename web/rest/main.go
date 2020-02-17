package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/health", HealthCheckHandler).Methods("GET")

	router.HandleFunc("/articles", HealthCheckHandler).Methods("GET")
	router.HandleFunc("/articles/{id}", HealthCheckHandler).Methods("GET")
	router.HandleFunc("/articles", HealthCheckHandler).Methods("POST")
	router.HandleFunc("/articles/{id}", HealthCheckHandler).Methods("PUT")
	router.HandleFunc("/articles/{id}", HealthCheckHandler).Methods("DELETE")
	router.HandleFunc("/articles/{id}/publish", HealthCheckHandler).
		Methods("PATCH")

	log.Println("api started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

type Article struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Body          string    `json:"body"`
	Published     bool      `json:"published"`
	CreatedDate   time.Time `json:"created_date"`
	UpdatedDate   time.Time `json:"updated_date"`
	PublishedDate time.Time `json:"published_date"`
	Authors       []string  `json:"authors`
}

func newArticle(Title, Body string, Published bool) Article {
	now := time.Now()
	a := Article{ID: uuid.New().String(), Title: Title, Body: Body,
		Published: Published, CreatedDate: now, PublishedDate: now}

	if Published {
		a.PublishedDate = now
	}

	return a
}

type HealthCheck struct {
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp"`
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(HealthCheck{Status: "OK",
		Timestamp: time.Now().UnixNano() / 1000000})
}

func ListArticleHandler(w http.ResponseWriter, r *http.Request) {

}

func GetArticleHandler(w http.ResponseWriter, r *http.Request) {

}

func AddArticleHandler(w http.ResponseWriter, r *http.Request) {

}

func UpdateArticleHandler(w http.ResponseWriter, r *http.Request) {

}

func PublishArticleHandler(w http.ResponseWriter, r *http.Request) {

}

func DeleteArticleHandler(w http.ResponseWriter, r *http.Request) {

}
