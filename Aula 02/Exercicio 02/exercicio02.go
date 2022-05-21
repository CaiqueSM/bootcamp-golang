package main

import (
	"errors"
	"fmt"
)

func calcularMedia(notas ...int) (float64, error) {
	soma := 0
	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("nota não pode ser negativa.")
		} else {
			soma += nota
		}
	}
	return float64(soma / len(notas)), nil
}

func main() {
	fmt.Println("A média das nota do aluno é:")
	fmt.Println(calcularMedia(10, 8, 9, 10))
}
