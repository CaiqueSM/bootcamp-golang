package main

import (
	"bufio"
	"math/rand"
	"os"
)

type Cliente struct {
	arquivo   int
	nome      string
	sobrenome string
	RG        string
	telefone  string
	endereco  string
}

func GerarID(semente int) int {
	return rand.Intn(semente)
}

func cadastrarCliente(cliente Cliente, clientes []Cliente) []Cliente {
	ID := GerarID(100)

	if ID == 0 {
		panic("ID não pode ser nulo")
	}

	cliente.arquivo = ID

	return append(clientes, cliente)
}

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

	novoCliente := Cliente{0, "C", "da SM", "98.765.456-0", "24910102020", "rua A, n 100, cidade das letras"}

	lerArquivo("customers.txt")

	cadastrarCliente(novoCliente, []Cliente{})

}
