package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// This approach works as a filter and allows access to the backend service
// (origin localhost:8080) only for the health check endpoint
// Test this proxy in front of `../basic_server`
// --
// See https://golang.org/pkg/net/http/httputil/#ReverseProxy for more details
// --
// This is a nice example https://blog.charmes.net/post/reverse-proxy-go/
func main() {

	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
	})

	log.Printf("Proxy starting on port %v\n", 8081)
	http.Handle("/health", proxy)
	http.ListenAndServe(":8081", nil)
}
