package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func minFunc(notas ...int) int {
	menor := notas[0]
	for _, nota := range notas {
		if menor > nota {
			menor = nota
		}
	}
	return menor
}

func averageFunc(notas ...int) int {
	soma := 0
	for _, nota := range notas {
		soma += nota
	}
	return soma / len(notas)
}

func maxFunc(notas ...int) int {
	maior := notas[0]
	for _, nota := range notas {
		if maior < nota {
			maior = nota
		}
	}
	return maior
}

func operation(operacao string) (func(notas ...int) int, error) {
	switch operacao {
	case minimum:
		return minFunc, nil
	case average:
		return averageFunc, nil
	case maximum:
		return maxFunc, nil
	}
	return nil, errors.New("operaçao nao reconhecida")
}

func main() {

	minValue, erro := operation(minimum)
	if erro != nil {
		fmt.Println(erro.Error())
	} else {
		fmt.Println("menor nota:", minValue(2, 3, 3, 4, 10, 2, 4, 5))
	}

	averageValue, erro := operation(average)
	if erro != nil {
		fmt.Println(erro.Error())
	} else {
		fmt.Println("média das notas:", averageValue(2, 3, 3, 4, 1, 2, 4, 5))
	}

	maxValue, erro := operation(maximum)
	if erro != nil {
		fmt.Println(erro.Error())
	} else {
		fmt.Println("maior nota:", maxValue(2, 3, 3, 4, 1, 2, 4, 5))
	}
}
