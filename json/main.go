package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type entryResponse struct {
	Message      string    `json:"message"`
	Time         time.Time `json:"timestamp,omitempty"`
	Code         int       `json:"code"`
	InternalCode int       `json:"-"`
}

type entryRequest struct {
	Title         string    `json:"title"`
	Slug          string    `json:"slug"`
	Body          string    `json:"body"`
	Published     bool      `json:"published"`
	PublishedDate time.Time `json:"published_date"`
	UpdatedDate   time.Time `json:"updated_date"`
}

func main() {
	port := 8080
	http.HandleFunc("/json1", simpleResponseEntryHandler)
	http.HandleFunc("/json2", improvedResponseEntryHandler)
	http.HandleFunc("/json3", simpleReqResEntryHandler)
	http.HandleFunc("/json4", improvedReqResEntryHandler)
	log.Printf("Listening at %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}

// $ curl -X GET http://localhost:8080/json1
func simpleResponseEntryHandler(w http.ResponseWriter, r *http.Request) {
	response := entryResponse{Message: "le message", Time: time.Now(), Code: 0}
	// data, err := json.MarshalIndent(response, "", "  ")
	data, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(data))
}

// Sends the data direct to the output stream without marshilling
// it to a temporary object. Using Encoder rather tan marshilling
// to a byte array is nearly 50% faster
//
// $ curl -X GET http://localhost:8080/json2
func improvedResponseEntryHandler(w http.ResponseWriter, r *http.Request) {
	response := entryResponse{Message: "le message", Time: time.Now(), Code: 0}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}

// $ curl -X POST http://localhost:8080/json3 -d
// '{"title": "foo", "updated_date": "2020-02-02T18:35:12.104585544-05:00", "published": false}'
func simpleReqResEntryHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading request body: %s", err.Error())
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var request entryRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Printf("error unmarshalling: %s", err.Error())
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	log.Printf("Request received: %s", string(body))

	response := entryResponse{Message: request.Title + " saved!", Time: time.Now()}
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

// $ curl -X POST http://localhost:8080/json4 -d
// '{"title": "foo", "updated_date": "2020-02-02T18:35:12.104585544-05:00", "published": false}'
func improvedReqResEntryHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var request entryRequest
	err := decoder.Decode(&request)

	if err != nil {
		log.Printf("error unmarshalling: %s", err.Error())
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	log.Printf("Request received: %s", request.UpdatedDate)

	response := entryResponse{Message: request.Title + " saved!", Time: time.Now()}
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
