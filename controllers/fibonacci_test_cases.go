package controllers

import (
	"net/http"
	"net/http/httptest"
)

type args struct {
	w http.ResponseWriter
	r *http.Request
}

// テストケースの構造体
var FibonacciTestCases = []struct {
	name           string
	args           args
	expectedStatus int
	expectedBody   string
}{
	//テストケース
	{
		name: "Valid request small number",
		args: args{
			w: httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet, "/fib?n=5", nil),
		},
		expectedStatus: http.StatusOK,
		expectedBody:   `{"result":5}`,
	},
	{
		name: "Valid request large number",
		args: args{
			w: httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet, "/fib?n=1000", nil),
		},
		expectedStatus: http.StatusOK,
		expectedBody:   `{"result":43466557686937456435688527675040625802564660517371780402481729089536555417949051890403879840079255169295922593080322634775209689623239873322471161642996440906533187938298969649928516003704476137795166849228875}`,
	},
	{
		name: "Valid request negative number",
		args: args{
			w: httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet, "/fib?n=10000000000", nil),
		},
		expectedStatus: http.StatusRequestTimeout,
		expectedBody:   `{"status":408,"message":"Request time out"}`,
	},
	{
		name: "Valid request negative number",
		args: args{
			w: httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet, "/fib?n=-1", nil),
		},
		expectedStatus: http.StatusBadRequest,
		expectedBody:   `{"status":400,"message":"Bad request"}`,
	},
	{
		name: "Valid request negative number",
		args: args{
			w: httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet, "/fib?n=abc", nil),
		},
		expectedStatus: http.StatusBadRequest,
		expectedBody:   `{"status":400,"message":"Bad request"}`,
	},
}
