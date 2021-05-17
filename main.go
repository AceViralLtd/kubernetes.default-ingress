package main

import (
	"fmt"
	"net/http"
	"time"
)

type RequestHandler struct {
}

// ServeHTTP is the entry method that the http lib will call to handler requests
func (r *RequestHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	_, _ = response.Write([]byte("Nope"))
}

// this almost seems like a waste of a service
func main() {
	server := &http.Server{
		Addr:           ":8081",
		Handler:        &RequestHandler{},
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   2 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println(server.ListenAndServe())
}
