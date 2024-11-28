package controllers

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFibonacciHandler(t *testing.T) {

	for _, tt := range FibonacciTestCases {
		t.Run(tt.name, func(t *testing.T) {
			//ハンドラー実行
			FibonacciHandler(tt.args.w, tt.args.r)

			//レスポンスの取得
			res := tt.args.w.(*httptest.ResponseRecorder)
			//ステータスコードのチェック
			if res.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, res.Code)
			}

			// レスポンスボディのチェック(予期しない改行が含まれるため、不要な改行は空白を除去する)
			bodyBytes, _ := io.ReadAll(res.Body)
			body := strings.TrimSpace(string(bodyBytes))

			if string(body) != tt.expectedBody {
				t.Errorf("expected body %q,got %q", tt.expectedBody, string(body))
			}
		})
	}
}
