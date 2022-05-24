package main

import (
	"bufio"
	"fmt"
	"os"
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

func imprimirArquivo(texto []string) {
	for _, each_ln := range texto {
		fmt.Println(each_ln)
	}
}
