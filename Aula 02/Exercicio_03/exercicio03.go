package main

import "fmt"

const (
	horasMensais        = 160
	salarioCategoriaA   = 3000
	salarioCategoriaB   = 1500
	salarioCategoriaC   = 1000
	adicionalCategoriaA = 50 / 100
	adicionalCategoriaB = 20 / 100
)

func calcularSalarioCategoriaA(horas int) float64 {
	if horas > horasMensais {
		return float64(salarioCategoriaA*horas +
			((salarioCategoriaA * (horas - horasMensais)) * adicionalCategoriaA))
	} else {
		return float64(salarioCategoriaA * horas)
	}
}

func calcularSalarioCategoriaB(horas int) float64 {
	if horas > horasMensais {
		return float64(salarioCategoriaB*horas +
			((salarioCategoriaB * (horas - horasMensais)) * adicionalCategoriaB))
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
