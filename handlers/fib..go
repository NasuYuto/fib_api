package handlers

import (
	"fmt"
	"math/big"
	"net/http"
	"strconv"
)

func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
	nStr := r.URL.Query().Get("n")
	n, err := strconv.Atoi(nStr)
	if err != nil || n < 0 {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
		return
	}

	result := fibnacci(n)
	fmt.Fprintf(w, "%d", result)
}

func fibnacci(n int) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}

	a, b := big.NewInt(0), big.NewInt(1)
	for i := 2; i <= n; i++ {
		a.Add(a, b)
		a, b = b, a
	}
	return b
}
