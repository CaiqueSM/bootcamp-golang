package main

import (
	"errors"
	"fmt"
)

const salarioMinimo = 15000

func VerificarSalario(valor int) error {

	if valor < salarioMinimo {
		return errors.New("erro: o salário digitado não está dentro do valor mínimo")
	}
	return nil
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
