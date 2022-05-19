package main

import "fmt"

var (
	temperatura int     = 18
	umidade     float32 = 43
	pressaoAtm  int     = 1013
)

func main() {
	fmt.Println("Temperartura:", temperatura)
	fmt.Println("Umidade:", umidade)
	fmt.Println("Pressão Atmosférica:", pressaoAtm)
}
