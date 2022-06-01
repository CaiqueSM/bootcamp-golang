package main

import "fmt"

const (
	horasMensais      = 160
	salarioCategoriaA = 3000
	salarioCategoriaB = 1500
	salarioCategoriaC = 1000
)

func calcularSalarioCategoriaA(horas int) float64 {
	if horas > horasMensais {
		return float64(horas) * salarioCategoriaA * 1.5
	} else {
		return float64(salarioCategoriaA * horas)
	}
}

func calcularSalarioCategoriaB(horas int) float64 {
	if horas > horasMensais {
		return float64(horas) * salarioCategoriaB * 1.2
	} else {
		return float64(salarioCategoriaB * horas)
	}
}

func calcularSalarioCategoriaC(horas int) float64 {
	return float64(salarioCategoriaC * horas)
}

func calcularSalario(categoria string) func(horas int) float64 {
	switch categoria {
	case "A":
		return calcularSalarioCategoriaA
	case "B":
		return calcularSalarioCategoriaB
	case "C":
		return calcularSalarioCategoriaC
	}

	return nil
}

func main() {
	categoriaA := calcularSalario("A")
	categoriaB := calcularSalario("B")
	categoriaC := calcularSalario("C")

	fmt.Println("Salário categoria A:", categoriaA(172))
	fmt.Println("Salário categoria B:", categoriaB(176))
	fmt.Println("Salário categoria :", categoriaC(162))
}
