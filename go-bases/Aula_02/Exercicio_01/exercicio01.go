package main

import "fmt"

func calcularImpostoSalario(Asalario, Ataxa float32) float32 {
	Ataxa /= 100
	return Asalario * Ataxa
}

func main() {

	var (
		imposto float32

		salario1 float32 = 50000
		salario2 float32 = 150000

		taxa1 float32 = 27
		taxa2 float32 = 10
	)

	imposto = calcularImpostoSalario(salario1, taxa1)
	fmt.Println("O imposto a pagar para o funcionario 1 é R$:", imposto)

	imposto = calcularImpostoSalario(salario2, taxa1)
	fmt.Println("O imposto a pagar para o funcionario 2 é R$:", imposto)

	imposto = calcularImpostoSalario(salario2, taxa2)
	fmt.Println("O imposto adicional a pagar para o funcionario 2 é R$:", imposto)

}
