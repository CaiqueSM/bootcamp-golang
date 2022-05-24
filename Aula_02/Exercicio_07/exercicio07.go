package main

import (
	"errors"
	"fmt"
)

const (
	pequeno = "pequeno"
	medio   = "medio"
	grande  = "grande"
)

type produto struct {
	tipo  string
	nome  string
	preco float64
}

func (p *produto) custoAdicionalProdutoPequeno() float64 {
	return p.preco
}

func (p *produto) custoAdicionalProdutoMedio() float64 {
	return p.preco * 1.03
}

func (p *produto) custoAdicionalProdutoGrande() float64 {
	return (p.preco * 1.06) + 2500
}

func (p *produto) CustoAdicional(tipo string) (func() float64, error) {
	switch tipo {
	case pequeno:
		return p.custoAdicionalProdutoPequeno, nil
	case medio:
		return p.custoAdicionalProdutoMedio, nil
	case grande:
		return p.custoAdicionalProdutoGrande, nil
	default:
		return nil, errors.New("tipo desconhecido")
	}
}

func (p produto) CalcularCusto() float64 {
	custoTotal, msg := p.CustoAdicional(p.tipo)
	if msg != nil {
		fmt.Println(msg.Error())
		return 0
	}
	return custoTotal()
}

type loja struct {
	listaProdutos []Produto
}

func (l loja) Total() float64 {
	soma := 0.0
	for _, produto := range l.listaProdutos {
		soma += produto.CalcularCusto()
	}
	return soma
}

func (l *loja) Adicionar(novoProduto Produto) {
	l.listaProdutos = append(l.listaProdutos, novoProduto)
}

type Produto interface {
	CalcularCusto() float64
}

type Ecommerce interface {
	Total() float64
	Adicionar(novoProduto Produto)
}

func NovoProduto(tipo, nome string, preco float64) Produto {
	p := produto{tipo, nome, preco}
	return p
}
func NovaLoja() Ecommerce {
	l := loja{}
	return &l
}

func main() {

	Loja01 := NovaLoja()

	produtoPequeno := NovoProduto(pequeno, "caneta", 1.50)
	produtoMedio := NovoProduto(medio, "cadeira", 180.90)
	produtoGrande := NovoProduto(grande, "mesa", 200.0)

	Loja01.Adicionar(produtoPequeno)
	Loja01.Adicionar(produtoMedio)
	Loja01.Adicionar(produtoGrande)

	fmt.Println("valor total da lista é:", Loja01.Total())

}
