package main

import "fmt"

const salarioMinimo = 15000

func VerificarSalario(valor int) error {

	if valor < salarioMinimo {
		return fmt.Errorf("erro: o mínimo tributável é 15.000 e o salário informado é: %d", valor)
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
