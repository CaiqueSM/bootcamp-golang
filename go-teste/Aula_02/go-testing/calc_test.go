package go_testing

import "testing"

func TestSub(t *testing.T){
	num1 := 5
	num2 := 2
	expectedResult := 3

	result := Sub(num1, num2)

	if result != expectedResult{
		t.Errorf("A função Sub() retornou o resultado = %v, mas o esperado é %v", result, expectedResult)
	}
}