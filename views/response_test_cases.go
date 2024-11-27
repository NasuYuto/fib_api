package views

import (
	"math/big"
	"net/http"
)

var WriteSuccessCases = []struct {
	name         string
	result       *big.Int
	expectedBody string
}{
	{
		name:         "Success with small number",
		result:       big.NewInt(5),
		expectedBody: `{"result":5}`,
	},
	{
		name:         "Success with large number",
		result:       big.NewInt(1000000000),
		expectedBody: `{"result":1000000000}`,
	},
}

var WriteErrorCases = []struct {
	name         string
	status       int
	message      string
	expectedBody string
}{
	{
		name:         "Bad request error",
		status:       http.StatusBadRequest,
		message:      "Bad request",
		expectedBody: `{"status":400,"message":"Bad request"}`,
	},
	{
		name:         "Request time out error",
		status:       http.StatusRequestTimeout,
		message:      "Request time out",
		expectedBody: `{"status":408,"message":"Request time out"}`,
	},
}
