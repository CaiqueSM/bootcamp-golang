package main

import "fmt"

var (
	idade        int
	salario      float32
	tempoServico uint8
)

func main() {
	fmt.Println("Entre com a idade:")
	fmt.Scanln(&idade)
	if idade < 22 {
		fmt.Println("Idade abaixo do permitido. Emprestimo recusado!")
	} else {
		fmt.Println("Entre com o tempo de serviço:")
		fmt.Scanln(&tempoServico)
		if tempoServico < 1 {
			fmt.Println("Tempo de serviço insuficiente. Emprestimo recusado!")
		} else {
			fmt.Println("Entre com o salário:")
			fmt.Scanln(&salario)
			if salario > 100000 {
				fmt.Println("Emprestimo concedido SEM cobrança de juros.")
			} else {
				fmt.Println("Emprestimo concedido COM cobrança de juros.")
			}
		}
	}
}
