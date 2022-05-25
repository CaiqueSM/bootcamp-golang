package main

import "fmt"

type Produto struct {
	Nome       string
	Preco      float64
	Quantidade int
}

type Usuario struct {
	Nome      string
	Sobrenome string
	Email     string
	Produtos  []Produto
}

func NovoProduto(nome string, preco float64) Produto {
	p := Produto{nome, preco, 0}
	return p
}

func AdicionarProduto(usuario *Usuario, produto *Produto, quantidade int) {
	produto.Quantidade = quantidade
	usuario.Produtos = append(usuario.Produtos, *produto)
}

func DeletarProdutos(usuario *Usuario) {
	usuario.Produtos = nil
}

func Detalhes(usuario *Usuario) {
	fmt.Println(usuario.Nome, usuario.Sobrenome)
	fmt.Println("Lista de produtos:")
	for _, produto := range usuario.Produtos {
		fmt.Println(produto.Nome)
	}
}

func main() {

	usuario1 := Usuario{
		Nome:      "c",
		Sobrenome: "sm",
		Email:     "csm@email.com",
		Produtos:  []Produto{},
	}

	produto1 := NovoProduto("caneta", 1.50)

	produto2 := NovoProduto("caderno", 12.90)

	produto3 := NovoProduto("calculadora", 5.90)

	AdicionarProduto(&usuario1, &produto1, 2)
	AdicionarProduto(&usuario1, &produto2, 1)

	Detalhes(&usuario1)

	DeletarProdutos(&usuario1)

	AdicionarProduto(&usuario1, &produto3, 1)

	Detalhes(&usuario1)
}
