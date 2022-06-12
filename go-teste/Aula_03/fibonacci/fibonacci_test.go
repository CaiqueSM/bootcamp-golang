package fibonacci_test

import (
	"testing"

	"github.com/CaiqueSM/bootcamp-golang.git/go-teste/Aula_03/fibonacci"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		arg  int
		want int
	}{{0, 0},{1, 1}}

	for i, d := range tests {
		got := fibonacci.Fibonacci(d.arg)
		if got != int64(d.want) {
			t.Errorf("Test[%d]: fibonacci(%d) returned %d, want %d",
				i, d.arg, got, d.want)
		}
	}
}
