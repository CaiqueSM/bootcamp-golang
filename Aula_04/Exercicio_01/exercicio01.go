package main


import "fmt"

const salarioMinimo = 15000

type ErroPersonalisado struct {
	msg string
}

func (e *ErroPersonalisado) Error() string {
	return fmt.Sprintf("%v", e.msg)
}

func VerificarSalario(valor int) error {

	if valor < salarioMinimo {
		return &ErroPersonalisado{msg: "Erro: O salário digitado não está dentro do valor mínimo"}
	}
	return nil
}

func main() {
	salario := 10000

	err := VerificarSalario(salario)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Necessário pagamento de imposto")
	}
}
