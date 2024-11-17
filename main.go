package main

import (
	"fib/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/fib", handlers.FibonacciHandler)

	log.Println("Starting server on :80")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
