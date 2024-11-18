package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestFibonacciHandler tests the FibonacciHandler function
func TestFibonacciHandler(t *testing.T) {
	tests := []struct {
		name           string
		queryParam     string
		expectedStatus int
		expectedResult string // JSONの "result" フィールド
		expectedError  string // JSONの "message" フィールド
	}{
		{"Valid small number", "n=5", http.StatusOK, "5", ""},
		{"Valid large number", "n=50", http.StatusOK, "12586269025", ""},
		{"Invalid negative number", "n=-1", http.StatusBadRequest, "", "Bad request"},
		{"Invalid non-integer input", "n=abc", http.StatusBadRequest, "", "Bad request"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/fib?"+tt.queryParam, nil)
			rec := httptest.NewRecorder()

			FibonacciHandler(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			// ステータスコードのチェック
			if res.StatusCode != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, res.StatusCode)
			}

			// レスポンスボディをJSONとして解析
			if tt.expectedStatus == http.StatusOK {
				var resp response
				if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}

				// "result" フィールドのチェック
				if resp.Result.String() != tt.expectedResult {
					t.Errorf("expected result %q, got %q", tt.expectedResult, resp.Result.String())
				}
			} else {
				var errResp errors
				if err := json.NewDecoder(res.Body).Decode(&errResp); err != nil {
					t.Fatalf("failed to decode error response: %v", err)
				}

				// "message" フィールドのチェック
				if errResp.Message != tt.expectedError {
					t.Errorf("expected error message %q, got %q", tt.expectedError, errResp.Message)
				}
			}
		})
	}
}

// TestFibnacci tests the fibnacci function directly
func TestFibnacci(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{"Fibonacci of 0", 0, "0"},
		{"Fibonacci of 1", 1, "1"},
		{"Fibonacci of 10", 10, "55"},
		{"Fibonacci of 50", 50, "12586269025"},
		{"Fibonacci of 100", 100, "354224848179261915075"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fibnacci(tt.input)
			if result.String() != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result.String())
			}
		})
	}
}
