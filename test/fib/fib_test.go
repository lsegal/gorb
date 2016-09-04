package fib_test

import (
	"testing"

	"github.com/lsegal/gorb/test/fib"
)

var tests = [][]int{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{10, 55},
}

func Test_Fibonacci_Fib(t *testing.T) {
	f := fib.Fibonacci{}
	for _, io := range tests {
		res := f.Fib(io[0])
		if res != io[1] {
			t.Fatalf("expected fib(%d) = %d, got %d", io[0], io[1], res)
		}
	}
}
