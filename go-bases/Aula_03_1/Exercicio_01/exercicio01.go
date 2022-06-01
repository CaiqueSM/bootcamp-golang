package main

import (
	"fmt"
)

type Usuario struct {
	nome      string
	sobrenome string
	idade     int
	email     string
	senha     string
}

func (u *Usuario) MudarNome(nome, sobrenome string) {
	u.nome = nome
	u.sobrenome = sobrenome
}

func (u *Usuario) MudarIdade(idade int) {
	u.idade = idade
}

func (u *Usuario) MudarEmail(email string) {
	u.email = email
}

func (u *Usuario) MudarSenha(senha string) {
	u.senha = senha
}

func (u *Usuario) Detalhes() {
	fmt.Println("Nome:", u.nome)
	fmt.Println("Sobrenome:", u.sobrenome)
	fmt.Println("Idade:", u.idade)
	fmt.Println("Email:", u.email)
}

func main() {
	usuario1 := Usuario{"Caíque", "da Silva Machado", 25, "generico@Email.com", "#####"}
	usuario1.Detalhes()
	fmt.Println("Atualizando...")
	usuario1.email = "generico.alternativo@Email.com"
	usuario1.Detalhes()
}
