package fibonacci_test

import (
	"testing"

	"github.com/CaiqueSM/bootcamp-golang.git/go-teste/Aula_03/fibonacci"
	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		arg  int
		want []int64
	}{{0, []int64{0}},{1, []int64{0, 1}},{10, []int64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}}}

	for i, d := range tests {
		got := fibonacci.Fibonacci(d.arg)
		if !assert.Equal(t, got, d.want) {
			t.Errorf("Test[%d]: fibonacci(%d) returned %d, want %d",
				i, d.arg, got, d.want)
		}
	}
}
