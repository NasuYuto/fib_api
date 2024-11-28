package views

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
)

type SuccessResponse struct {
	Result *big.Int `json:"result"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func WriteSuccessResponse(w http.ResponseWriter, result *big.Int) {
	//レスポンスの内容がJSONであることを明示
	w.Header().Set("Content-Type", "application/json")
	//レスポンスのステータスコードを伝える
	w.WriteHeader(http.StatusOK)

	//構造体に格納
	resp := SuccessResponse{Result: result}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		fmt.Printf("Error encoding success response: %v\n", err)
	}
}

func WriteErrorResponse(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "applocation/json")
	w.WriteHeader(status)

	errResp := ErrorResponse{
		Status:  status,
		Message: message,
	}
	if err := json.NewEncoder(w).Encode(errResp); err != nil {
		fmt.Printf("Error encoding error response: %v\n", err)
	}
}
