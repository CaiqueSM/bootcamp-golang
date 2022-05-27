package main

import (
	"bufio"
	"fmt"
	"os"
)

func lerArquivo(caminho string) []string {
	arquivo, err := os.Open(caminho)

	if err != nil {
		panic("o arquivo indicado não foi encontrado ou está danificado")
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

func main() {
	fmt.Println("Iniciando... ")
	lerArquivo("customers.txt")
	fmt.Println("Fim")
}
