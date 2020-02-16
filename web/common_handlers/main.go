package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {

	// handleFileServer()
	// handleNotFound()
	// handleRedirect()
	// handleStripPrefix()
	// handleTimeout()

	port := 8080
	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func handleRedirect() {
	http.Handle("/search", http.RedirectHandler("https://google.com",
		http.StatusMovedPermanently))
}

func handleFileServer() {
	http.Handle("/", http.FileServer(http.Dir("./data")))
}

func handleNotFound() {
	http.Handle("/foo", http.NotFoundHandler())
}

func handleStripPrefix() {
	http.Handle("/data", http.StripPrefix("/data",
		http.FileServer(http.Dir("./data"))))
}

func handleTimeout() {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 3)
		io.WriteString(w, "hey!")
	})
	http.Handle("/hey", http.TimeoutHandler(h, time.Second*2, "server timeout"))
}
