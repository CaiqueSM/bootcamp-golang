package main

import (
	"fmt"
	"os"
	"strings"
)

func escreverArquivo() {
	var (
		ID         int
		preco      float64
		quantidade int
	)

	cabecalho := "ID\t Preço\t Quantidade\n"
	resposta := ""

	arquivo, err := os.Create("produto.txt")
	arquivo.Write([]byte(cabecalho))

	if err != nil {
		fmt.Println(err.Error())
	}

	for {
		fmt.Println("Adicionar novo produto: S/N")
		fmt.Scanln(&resposta)
		if strings.ToUpper(resposta) == "S" {
			fmt.Println("Informe o ID do produto:")
			fmt.Scanln(&ID)
			fmt.Println("Informe o preço do produto:")
			fmt.Scanln(&preco)
			fmt.Println("Informe a quantidade:")
			fmt.Scanln(&quantidade)

			linha := []byte(fmt.Sprint(ID, "\t", preco, "\t", quantidade, ";\n"))
			arquivo.Write(linha)

			if err != nil {
				fmt.Println(err.Error())
			}

		} else if strings.ToUpper(resposta) == "N" {
			return
		} else {
			continue
		}
	}
}

func main() {
	escreverArquivo()
	fmt.Println("Arquivo escrito!")
}
