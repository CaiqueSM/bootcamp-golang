package main

import (
	"errors"
	"fmt"
)

const (
	salarioMinimo = 15000
	salarioTributavel = 20000
	valorHora     = 2000
)

func VerificarSalario(valor int) error {

	if valor < salarioMinimo {
		return fmt.Errorf("erro: o mínimo tributável é 15.000 e o salário informado é: %d", valor)
	}
	return nil
}

func CalcularSalario(valorHora int, horasTrabalhadas int)(int, error) {
	salario := valorHora * horasTrabalhadas

	if horasTrabalhadas < 0{
		err := errors.New("erro: as horas trabalhadas nao pode ser negativa")
		return 0, err
	}

	if horasTrabalhadas < 80{
		err := errors.New("erro: o trabalhador não pode ter trabalhado menos de 80 horas por mês")
		return 0, err
	}

	if salario >= salarioTributavel{
		salario = int(float64(salario) * float64(0.9))
	}
	return salario, nil
}

func main() {
	salario := 10000

	err := VerificarSalario(salario)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Necessário pagamento de imposto")
	}
}
