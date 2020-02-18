package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func main() {

	LoadMock()

	router := mux.NewRouter()

	router.HandleFunc("/health", HealthCheckHandler).Methods("GET")

	router.HandleFunc("/articles", ListArticleHandler).Methods("GET")
	router.HandleFunc("/articles/{id}", GetArticleHandler).Methods("GET")
	router.HandleFunc("/articles", AddArticleHandler).Methods("POST")
	router.HandleFunc("/articles/{id}", UpdateArticleHandler).Methods("PUT")
	router.HandleFunc("/articles/{id}", DeleteArticleHandler).Methods("DELETE")
	router.HandleFunc("/articles/{id}", PublishArticleHandler).Methods("PATCH")

	log.Println("api started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

type Article struct {
	ID            string     `json:"id"`
	Title         string     `json:"title"`
	Body          string     `json:"body"`
	Published     bool       `json:"published"`
	CreatedDate   time.Time  `json:"created_date"`
	UpdatedDate   time.Time  `json:"updated_date"`
	PublishedDate *time.Time `json:"published_date,omitempty"`
}

func newArticle(Title, Body string, Published bool) Article {
	now := time.Now()
	a := Article{ID: uuid.New().String(), Title: Title, Body: Body,
		Published: Published, CreatedDate: now, UpdatedDate: now}

	if Published {
		a.PublishedDate = &now
	}

	return a
}

type HealthCheck struct {
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp"`
}

type ApiError struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	Code      int    `json:"code"`
	Timestamp int64  `json:"timestamp"`
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(HealthCheck{Status: "OK",
		Timestamp: epochMilis()})
}

func ListArticleHandler(w http.ResponseWriter, r *http.Request) {
	result, err := mockList()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	writeResponse(w, http.StatusOK, result)
}

func GetArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	exists, err := mockExists(id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	if !exists {
		writeError(w, http.StatusNotFound, errors.New("not found"))
		return
	}

	a, err := mockGet(id)

	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeResponse(w, http.StatusOK, a)
}

func AddArticleHandler(w http.ResponseWriter, r *http.Request) {
	var a Article
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&a); err != nil {
		writeError(w, http.StatusBadRequest, errors.New("invalid request"))
		return
	}

	if createdArticle, err := mockSave(a); err == nil {
		writeResponse(w, http.StatusCreated, createdArticle)
	} else {
		writeError(w, http.StatusInternalServerError, err)
	}
}

func UpdateArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	exists, err := mockExists(id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	if !exists {
		writeError(w, http.StatusNotFound, errors.New("not found"))
		return
	}

	var a Article
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&a); err != nil {
		writeError(w, http.StatusBadRequest, errors.New("invalid request"))
		return
	}

	if _, err := mockSave(a); err == nil {
		writeResponse(w, http.StatusNoContent, nil)
	} else {
		writeError(w, http.StatusInternalServerError, err)
	}
}

func PublishArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	exists, err := mockExists(id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	if !exists {
		writeError(w, http.StatusNotFound, errors.New("not found"))
		return
	}

	a, err := mockGet(id)

	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	a.Published = true
	now := time.Now()
	a.PublishedDate = &now

	if _, err := mockSave(*a); err == nil {
		writeResponse(w, http.StatusNoContent, nil)
	} else {
		writeError(w, http.StatusInternalServerError, err)
	}
}

func DeleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	exists, err := mockExists(id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	if !exists {
		writeError(w, http.StatusNotFound, errors.New("not found"))
		return
	}

	if err := mockDelete(id); err == nil {
		writeResponse(w, http.StatusNoContent, nil)
	} else {
		writeError(w, http.StatusInternalServerError, err)
	}
}

func writeResponse(w http.ResponseWriter, code int, payload interface{}) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		log.Print(err.Error())
	}
}

func writeError(w http.ResponseWriter, code int, err error) {
	// if code is 5XX we shouldn't send details to the client in production
	// skipping this check here
	e := ApiError{ID: uuid.New().String(), Code: code, Message: err.Error(),
		Timestamp: epochMilis()}
	log.Printf("%s - %s", e.ID, err.Error())
	writeResponse(w, code, e)
}

func epochMilis() int64 {
	return time.Now().UnixNano() / 1000000
}

//////////
// Mock //
//////////

var MockArticles map[string]Article

func LoadMock() {
	MockArticles = make(map[string]Article)
	for i := 0; i < 10; i++ {
		a :=
			newArticle(fmt.Sprintf("Title %d", i), fmt.Sprintf("Body %d", i),
				false)
		MockArticles[a.ID] = a
	}
}

func mockList() ([]Article, error) {
	result := []Article{}
	for _, v := range MockArticles {
		result = append(result, v)
	}
	return result, nil
}

func mockExists(ID string) (bool, error) {
	if _, ok := MockArticles[ID]; ok {
		return true, nil
	}
	return false, nil
}

func mockGet(ID string) (*Article, error) {
	if a, ok := MockArticles[ID]; ok {
		return &a, nil
	}
	return nil, errors.New("no result error")
}

func mockSave(a Article) (*Article, error) {
	MockArticles[a.ID] = a
	return &a, nil
}

func mockDelete(ID string) error {
	delete(MockArticles, ID)
	return nil
}
