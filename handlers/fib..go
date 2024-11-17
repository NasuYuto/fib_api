package handlers

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
)

type response struct {
	Result *big.Int `json:"result"`
}

type errors struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータ "n" を取得
	nStr := r.URL.Query().Get("n")
	n, err := strconv.Atoi(nStr)
	if err != nil || n < 0 {
		// エラーをJSON形式で返す
		writeErrorResponse(w, http.StatusBadRequest, "Bad request")
		return
	}

	// フィボナッチ計算結果を構造体に格納
	result := response{
		Result: fibnacci(n),
	}

	// 結果をJSON形式に変換してレスポンスに書き込み
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		// JSONエンコードエラーの場合もエラーレスポンスを返す
		writeErrorResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}
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

func writeErrorResponse(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status) // ステータスコードを設定

	// エラー構造体を作成
	errResp := errors{
		Status:  status,
		Message: message,
	}

	// JSON形式でレスポンスを送信
	if err := json.NewEncoder(w).Encode(errResp); err != nil {
		fmt.Printf("Error encoding error response: %v\n", err)
	}
}
