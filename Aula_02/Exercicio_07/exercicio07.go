package main

type produto struct {
	tipo  string
	nome  string
	preco float64
}

type loja struct {
	listaProdutos produto
}

type Produto interface {
	CalcularCusto()
}

type Ecommerce interface {
	Total()
	Adicionar()
	novoProduto(tipo, nome string, preco float64) Produto
	novaLoja() Ecommerce
}
