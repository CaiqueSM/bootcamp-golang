package main

import (
	"fmt"
	"github.com/CaiqueSM/bootcamp-golang.git/go-teste/Aula_03/fibonacci"
)

func main(){
	for i := 0; i <= 10 ;i++{
		fmt.Println(fibonacci.Fibonacci(i))
	}	
}