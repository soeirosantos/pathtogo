package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
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
	json.NewDecoder(rr.Body).Decode(&hc)

	expected := "OK"

	if hc.Status != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			hc.Status, expected)
	}
}
