package controllers

import (
	"context"
	"fib/models"
	"fib/views"
	"net/http"
	"strconv"
	"time"
)

func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	nStr := r.URL.Query().Get("n")
	n, err := strconv.Atoi(nStr)
	if err != nil || n <= 0 {
		// エラーをJSON形式で返す
		views.WriteErrorResponse(w, http.StatusBadRequest, "Bad request")
		return
	}

	resultChan := make(chan *models.FibonacciResult)
	go func() {
		result := models.CalculateFibonacci(n)
		select {
		case resultChan <- &models.FibonacciResult{Value: result}:
		case <-ctx.Done():
		}
	}()

	select {
	case <-ctx.Done():
		//タイムアウトの時のレスポンス
		views.WriteErrorResponse(w, http.StatusRequestTimeout, "Request time out")
	case res := <-resultChan:
		//正常に結果が返ってきた場合
		views.WriteSuccessResponse(w, res.Value)
	}
}
