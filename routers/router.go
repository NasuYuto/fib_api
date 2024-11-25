package routers

import (
	"fib/controllers"
	"net/http"
)

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/fib", controllers.FibonacciHandler)

	return router
}
