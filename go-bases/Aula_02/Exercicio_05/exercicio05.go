package main

import (
	"errors"
	"fmt"
)

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

const (
	qntAlimentoDog       = 10
	qntAlimentoCat       = 5
	qntAlimentoHamster   = float64(250) / float64(1000)
	qntAlimentoTarantula = float64(150) / float64(1000)
)

func quantidadeAlimentoDog(quantidade float64) float64 {
	return qntAlimentoDog * quantidade
}

func quantidadeAlimentoCat(quantidade float64) float64 {
	return qntAlimentoCat * quantidade
}

func quantidadeAlimentoHamster(quantidade float64) float64 {
	return qntAlimentoHamster * quantidade
}

func quantidadeAlimentoTarantula(quantidade float64) float64 {
	return qntAlimentoTarantula * quantidade
}

func Animal(animal string) (func(quantidade float64) float64, error) {
	switch animal {
	case dog:
		return quantidadeAlimentoDog, nil
	case cat:
		return quantidadeAlimentoCat, nil
	case hamster:
		return quantidadeAlimentoHamster, nil
	case tarantula:
		return quantidadeAlimentoTarantula, nil
	}

	return nil, errors.New("Animal nao encontrado")
}

func main() {
	var amount float64 = 0

	animalDog, msg := Animal(dog)
	if msg != nil {
		fmt.Println(msg.Error())
	} else {
		amount += animalDog(5)
	}

	animalCat, msg := Animal(cat)
	if msg != nil {
		fmt.Println(msg.Error())
	} else {
		amount += animalCat(8)
	}

	animalhamster, msg := Animal(hamster)
	if msg != nil {
		fmt.Println(msg.Error())
	} else {
		amount += animalhamster(3)
	}

	animaltarantula, msg := Animal(tarantula)
	if msg != nil {
		fmt.Println(msg.Error())
	} else {
		amount += animaltarantula(1)
	}
	fmt.Println("Qunatidade de alimento:", amount)
}
