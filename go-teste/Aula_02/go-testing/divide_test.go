package go_testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivide(t *testing.T){
	num := 10
	den1 := 2
	den2 := 0
	expectedResult := 5
	Result, err1 := Divide(num, den1)
	_, err2 := Divide(num, den2)

	assert.Equal(t, expectedResult, Result)
	assert.Nil(t, err1)
	assert.NotNil(t, err2)
}