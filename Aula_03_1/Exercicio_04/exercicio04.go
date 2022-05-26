package main

import (
	"fmt"
	"math/rand"
	"time"
)

func MetodoInsercao(vetor []int) []int {
	i := 0
	j := 1
	aux := 0
	n := len(vetor)

	for j < n {
		aux = vetor[j]
		i = j - 1
		for (i >= 0) && (vetor[i] > aux) {
			vetor[i+1] = vetor[i]
			i--
		}
		vetor[i+1] = aux
		j++
	}
	return vetor
}

func MetodoSelecao(vetor []int) []int {
	var min, aux int
	n := len(vetor)

	for i := 0; i < n-1; i++ {
		min = 1
		for j := (i + 1); j < n; j++ {
			if vetor[j] < vetor[min] {
				min = j
			}
		}
		if vetor[i] != vetor[min] {
			aux = vetor[i]
			vetor[i] = vetor[min]
			vetor[min] = aux
		}
	}
	return vetor
}

func Agrupar(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func MetodoGrupo(items []int) []int {
	if len(items) < 2 {
		return items
	}
	first := MetodoGrupo(items[:len(items)/2])
	second := MetodoGrupo(items[len(items)/2:])
	return Agrupar(first, second)
}

func ClassificarTempo(inicio time.Time, nome string) {
	final := time.Since(inicio)
	fmt.Printf("%s levou %d ms \n", nome, final.Milliseconds())
}

func ClassificarMetodoGrupo() {

	listaCem := rand.Perm(100)
	listaMil := rand.Perm(1000)
	listaDezMil := rand.Perm(10000)

	defer ClassificarTempo(time.Now(), "Metodo grupo:")

	MetodoGrupo(listaCem)
	MetodoGrupo(listaMil)
	MetodoGrupo(listaDezMil)
}

func ClassificarMetodoIsercao() {

	listaCem := rand.Perm(100)
	listaMil := rand.Perm(1000)
	listaDezMil := rand.Perm(10000)

	defer ClassificarTempo(time.Now(), "Metodo insercao:")

	MetodoInsercao(listaCem)
	MetodoInsercao(listaMil)
	MetodoInsercao(listaDezMil)
}

func ClassificarMetodoSelecao() {

	listaCem := rand.Perm(100)
	listaMil := rand.Perm(1000)
	listaDezMil := rand.Perm(10000)

	defer ClassificarTempo(time.Now(), "Metodo selecao:")

	MetodoSelecao(listaCem)
	MetodoSelecao(listaMil)
	MetodoSelecao(listaDezMil)
}

func main() {

	ClassificarMetodoGrupo()
	ClassificarMetodoIsercao()
	ClassificarMetodoSelecao()
}
