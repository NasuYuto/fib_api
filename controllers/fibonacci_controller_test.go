package controllers

import (
	"io/ioutil"
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

			// レスポンスボディのチェック
			bodyBytes, _ := ioutil.ReadAll(res.Body)
			body := strings.TrimSpace(string(bodyBytes))

			if string(body) != tt.expectedBody {
				t.Errorf("expected body %q,got %q", tt.expectedBody, string(body))
			}
		})
	}
}
