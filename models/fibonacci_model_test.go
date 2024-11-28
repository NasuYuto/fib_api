package models

import (
	"reflect"
	"testing"
)

func TestCalculateFibonacci(t *testing.T) {

	for _, tt := range FibonacciTestCases {
		t.Run(tt.name, func(t *testing.T) {
			//DeepEqual(ディープイコール)の引数値が不一致だった場合trueになりerror処理
			if got := CalculateFibonacci(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateFibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
