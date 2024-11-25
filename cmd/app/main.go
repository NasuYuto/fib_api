package main

import (
	"fib/routers"
	"log"
	"net/http"
)

func main() {
	router := routers.NewRouter()
	log.Println("Starting HTTP server on :80")
	if err := http.ListenAndServe(":80", router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
