package main

import "fmt"

var numeroMes int
var meses = []string{"Janeiro", "Fevereiro", "Março",
	"Abril", "Maio", "Junho", "Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro"}

func main() {
	fmt.Println("Digite o número do mês:")
	fmt.Scanln(&numeroMes)
	fmt.Println("Você escolheu", meses[numeroMes-1])
}
