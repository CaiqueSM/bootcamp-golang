package main

import (
	"errors"
	"fmt"
)

const (
	salarioTributavel = 20000
	valorHora         = 100
)

func CalcularSalario(valorHora int, horasTrabalhadas int) (int, error) {
	salario := valorHora * horasTrabalhadas

	if horasTrabalhadas < 80 {
		err := errors.New("erro: o trabalhador não pode ter trabalhado menos de 80 horas por mês")
		return 0, err
	}

	if salario >= salarioTributavel {
		salario = int(float64(salario) * float64(0.9))
	}
	return salario, nil
}

func main() {

	salario, err1 := CalcularSalario(valorHora, 79)

	err2 := fmt.Errorf("err1: %w", err1)

	if err1 != nil {
		fmt.Println(errors.Unwrap(err2))
		fmt.Println(errors.Unwrap(err1))
	}

	fmt.Println("Salário após as dedução de imposto", salario)

}
