package main

import "fmt"

func Fibonacci(n int) int64 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main(){
	for i := 0; i < 10 ;i++{
		fmt.Println(Fibonacci(i))
	}	
}