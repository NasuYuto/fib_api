package models

import (
	"reflect"
	"testing"
)

func TestCalculateFibonacci(t *testing.T) {

	for _, tt := range FibonacciTestCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateFibonacci(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateFibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
