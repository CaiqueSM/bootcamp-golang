package main

import (
	"math/rand"
	"strings"
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
	ID := rand.Intn(semente)
	return ID
}

func cadastrarCliente(cliente Cliente){
	ID := GerarID

	if ID == 0{
		panic("ID não pode ser nulo")
	}

	cliente.arquivo := ID
	
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

func recuperarDados(texto []string)[]Cliente{
	var clientes []Cliente
	for _, linha := range texto{
		info := strings.Split(linha, " ")
		clientes = append(clientes, Cliente{info[0], info[1], info[2], info[3], info[4], info[5] })
	}
	return clientes
}

func Existe(RG string, clientes []Cliente)bool{
	for _,cliente := range clientes{
		if RG == cliente.RG{
			return true
		}
	}
	return false
}

func main(){
	
}

