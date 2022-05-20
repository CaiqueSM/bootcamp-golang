package main

import "fmt"

var palavra string

func main() {
	fmt.Scanf("%s", &palavra)
	fmt.Println("O tamanho da palavra é:", len(palavra))
	fmt.Println("Soletrando:")
	for i := 0; i < len(palavra); i++ {
		fmt.Printf("%c\n", palavra[i])
	}
}
