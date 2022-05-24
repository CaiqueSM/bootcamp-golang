package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func lerArquivo(caminho string) []string {
	arquivo, err := os.Open(caminho)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	leitor := bufio.NewScanner(arquivo)
	leitor.Split(bufio.ScanLines)

	var texto []string

	for leitor.Scan() {
		texto = append(texto, leitor.Text())
	}

	arquivo.Close()

	return texto
}

func calculaPrecoTotal(texto []string) float64 {
	soma := 0.0
	for i := 1; i < len(texto); i++ {
		linha := texto[i]
		valores := strings.Split(linha, "\t")
		preco, _ := strconv.ParseFloat(valores[1], 64)
		quantidade, err := strconv.ParseFloat(valores[2], 64)
		if err != nil {
			fmt.Println(err.Error())
			return 0
		}

		soma += float64(preco * quantidade)
	}
	return soma
}

func imprimirArquivo(texto []string) {
	for _, linha := range texto {
		fmt.Println(linha)
	}
	fmt.Println("\t", calculaPrecoTotal(texto))
}

func main() {
	t := lerArquivo("/Users/caimachado/Documents/Github/bootcamp-golang/Aula_03/Exercicio_01/produto.csv")
	imprimirArquivo(t)
}
