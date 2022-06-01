package main

import "fmt"

type Aluno struct {
	primeiroNome   string
	sobrenome      string
	rg             string
	dataDeAdmissao string
}

func (a *Aluno) Detalhes() {
	fmt.Println("Nome:", a.primeiroNome)
	fmt.Println("Sobrenome:", a.sobrenome)
	fmt.Println("RG:", a.rg)
	fmt.Println("Data admissão:", a.dataDeAdmissao)
}

func main() {
	aluno1 := Aluno{"Paula", "Monteiro", "24.123.456-00", "16-05-2022"}
	aluno2 := Aluno{"Fernando", "Silva", "24.098.765.44", "16-05-2021"}

	aluno1.Detalhes()
	aluno2.Detalhes()
}
