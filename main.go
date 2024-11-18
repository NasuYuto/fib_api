package main

import (
	"fib/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/fib", handlers.FibonacciHandler)

	log.Println("Starting HTTP server on :80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
