package main

import "fmt"

type Produto struct {
	nome       string
	preco      float64
	quantidade int
}

type Servico struct {
	nome    string
	preco   float64
	minutos int
}

type Manutencao struct {
	nome  string
	preco float64
}

func SomarProdutos(produtos []Produto, canal chan float64) float64 {
	soma := 0.0
	for _, produto := range produtos {
		soma += produto.preco * float64(produto.quantidade)
	}
	canal <- soma
	return soma
}

func SomarServicos(servicos []Servico, canal chan float64) float64 {

	const horaMediaTrabalho = 30

	soma := 0.0
	for _, servico := range servicos {
		if servico.minutos < horaMediaTrabalho {
			soma += servico.preco * horaMediaTrabalho
		} else {
			soma += servico.preco * float64(servico.minutos)
		}
	}
	canal <- soma
	return soma
}

func SomarManutencao(manutencoes []Manutencao, canal chan float64) float64 {
	soma := 0.0
	for _, manutencao := range manutencoes {
		soma += manutencao.preco
	}
	canal <- soma
	return soma
}

func main() {

	produtos := []Produto{{"roteador", 120.0, 1}, {"fibra otica", 50.0, 10},
		{"conectores", 20.0, 30}}

	servicos := []Servico{{"instalacao", 20.0, 60}, {"instalacao", 20.0, 120},
		{"instalacao", 20.0, 15}, {"instalacao", 20.0, 45}}

	manutencoes := []Manutencao{{"fibra otica", 5.99}, {"roteado", 8.99},
		{"conectores", 5.49}, {"fibra otica", 5.99}}

	canal := make(chan float64)
	precoTotal := 0.0

	go SomarProdutos(produtos, canal)
	precoTotal += <-canal

	go SomarServicos(servicos, canal)
	precoTotal += <-canal

	go SomarManutencao(manutencoes, canal)
	precoTotal += <-canal

	fmt.Println("O preco total das atividades é:", precoTotal)
}
