package views

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWriteSuccessResponse(t *testing.T) {

	for _, tt := range WriteSuccessCases {
		t.Run(tt.name, func(t *testing.T) {
			//HTTPレスポンスを模倣するためのオブジェクト
			recorder := httptest.NewRecorder()
			//tt.resultを使ってレスポンスを生成
			WriteSuccessResponse(recorder, tt.result)

			res := recorder.Result()
			//後にリソースを解放
			defer res.Body.Close()

			// レスポンスボディのチェック
			if strings.TrimSpace(recorder.Body.String()) != tt.expectedBody {
				t.Errorf("unexpected response body: got %v, want %v", recorder.Body.String(), tt.expectedBody)
			}

			// ステータスコードのチェック
			if res.StatusCode != http.StatusOK {
				t.Errorf("unexpected status code: got %v, want %v", res.StatusCode, http.StatusOK)
			}
		})
	}
}

func TestWriteErrorResponse(t *testing.T) {

	for _, tt := range WriteErrorCases {
		t.Run(tt.name, func(t *testing.T) {

			recorder := httptest.NewRecorder()
			//tt.statusとtt.messageを使ってレスポンスを生成
			WriteErrorResponse(recorder, tt.status, tt.message)

			res := recorder.Result()
			defer res.Body.Close()

			// レスポンスボディのチェック
			if strings.TrimSpace(recorder.Body.String()) != tt.expectedBody {
				t.Errorf("unexpected response body: got %v, want %v", recorder.Body.String(), tt.expectedBody)
			}

			// ステータスコードのチェック
			if res.StatusCode != tt.status {
				t.Errorf("unexpected status code: got %v, want %v", res.StatusCode, tt.status)
			}
		})
	}
}
