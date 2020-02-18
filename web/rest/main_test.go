package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func Test_healthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var hc HealthCheck
	err = json.NewDecoder(rr.Body).Decode(&hc)
	if err != nil {
		t.Error(err)
		return
	}

	expected := "OK"

	if hc.Status != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			hc.Status, expected)
	}
}

func TestListArticleHandler(t *testing.T) {
	LoadMock()
	req, _ := http.NewRequest("GET", "/articles", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListArticleHandler)

	handler.ServeHTTP(rr, req)

	var articles []Article
	err := json.NewDecoder(rr.Body).Decode(&articles)
	if err != nil {
		t.Error(err)
		return
	}

	if len(articles) == 0 {
		t.Error("handler returned a wrong number of results body")
	}
}

func TestGetArticleHandler(t *testing.T) {
	LoadMock()

	var id string
	var expectedArticle Article
	for k, a := range MockArticles {
		id = k
		expectedArticle = a
		break
	}

	req, _ := http.NewRequest("GET", fmt.Sprintf("/articles/%s", id), nil)

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/articles/{id}", GetArticleHandler)

	router.ServeHTTP(rr, req)

	var article Article
	err := json.NewDecoder(rr.Body).Decode(&article)
	if err != nil {
		t.Error(err)
		return
	}

	if article.ID != expectedArticle.ID {
		t.Errorf("expected id %s got %s", expectedArticle.ID, article.ID)
	}
}

func TestAddArticleHandler(t *testing.T) {
	LoadMock()

	payload :=
		[]byte(`{"title": "Test Article", "body": "Some body", "published": false}`)

	req, _ := http.NewRequest("POST", "/articles", bytes.NewBuffer(payload))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddArticleHandler)

	handler.ServeHTTP(rr, req)

	var article Article
	err := json.NewDecoder(rr.Body).Decode(&article)
	if err != nil {
		t.Error(err)
		return
	}

	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status code %d got %d", http.StatusCreated, rr.Code)
		return
	}

	expected := "Test Article"

	if article.Title != expected {
		t.Errorf("Expected title %s got %s", expected, article.Title)
	}
}

func TestUpdateArticleHandler(t *testing.T) {
	LoadMock()

	var id string
	var originalArticle Article
	for k, a := range MockArticles {
		id = k
		originalArticle = a
		break
	}

	originalArticle.Title = "New Title"

	payload, _ := json.Marshal(originalArticle)

	req, _ := http.NewRequest("PUT", fmt.Sprintf("/articles/%s", id),
		bytes.NewBuffer(payload))

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/articles/{id}", UpdateArticleHandler)

	router.ServeHTTP(rr, req)

	changedArticle := MockArticles[id]

	if rr.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d got %d", http.StatusNoContent, rr.Code)
		return
	}

	if changedArticle.ID != originalArticle.ID {
		t.Errorf("expected id %s got %s", originalArticle.ID, changedArticle.ID)
		return
	}

	if changedArticle.Title != originalArticle.Title {
		t.Errorf("expected title %s got %s", originalArticle.ID, changedArticle.ID)
		return
	}

	if changedArticle.Body != originalArticle.Body {
		t.Errorf("expected body %s got %s", originalArticle.Body, changedArticle.Body)
	}
}

func TestPublishArticleHandler(t *testing.T) {
	LoadMock()

	var id string
	var originalArticle Article
	for k, a := range MockArticles {
		id = k
		originalArticle = a
		break
	}

	req, _ := http.NewRequest("PATCH", fmt.Sprintf("/articles/%s", id), nil)

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/articles/{id}", PublishArticleHandler)

	router.ServeHTTP(rr, req)

	publishedArticle := MockArticles[id]

	if rr.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d got %d", http.StatusNoContent, rr.Code)
		return
	}

	if publishedArticle.ID != originalArticle.ID {
		t.Errorf("expected id %s got %s", originalArticle.ID, publishedArticle.ID)
		return
	}

	if publishedArticle.Title != originalArticle.Title {
		t.Errorf("expected title %s got %s", originalArticle.ID, publishedArticle.ID)
		return
	}

	if !publishedArticle.Published {
		t.Errorf("expected published %v got %v", true, publishedArticle.Published)
	}
}
