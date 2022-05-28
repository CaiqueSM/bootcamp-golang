package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const semente = 1000

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
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	ID := GerarID(semente)

	if ID == 0 {
		panic("ID não pode ser nulo")
	}

	cliente.arquivo = ID

	return append(clientes, cliente)
}

func lerArquivo(caminho string) []string {

	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

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

func recuperarDados(texto []string) []Cliente {
	var clientes []Cliente
	for _, linha := range texto {
		info := strings.Split(linha, " ")
		id, _ := strconv.Atoi(info[0])
		clientes = append(clientes, Cliente{id, info[1], info[2], info[3], info[4], info[5]})
	}
	return clientes
}

func Existe(RG string, clientes []Cliente) bool {
	for _, cliente := range clientes {
		if RG == cliente.RG {
			return true
		}
	}
	return false
}

func imprimirClientes(clientes []Cliente) {
	for _, cliente := range clientes {
		fmt.Println(cliente.arquivo, cliente.nome, cliente.sobrenome, cliente.RG, cliente.telefone, cliente.endereco)
	}
}

func main() {

	novoCliente := Cliente{0, "C", "da SM", "98.765.456-0", "24910102020", "rua A, n 100, cidade das letras"}

	arquivo := lerArquivo("customers.txt")
	Dados := recuperarDados(arquivo)

	if !Existe(novoCliente.RG, Dados) {
		Dados = cadastrarCliente(novoCliente, Dados)
	}
	imprimirClientes(Dados)
}
