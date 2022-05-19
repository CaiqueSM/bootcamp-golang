package main

import "fmt"

var (
	temperatura string = "18.0 °C"
	umidade     string = "43%"
	pressaoAtm  string = "1013 hPa"
)

func main() {
	fmt.Println("Temperartura:", temperatura)
	fmt.Println("Umidade:", umidade)
	fmt.Println("Pressão Atmosférica:", pressaoAtm)
}
