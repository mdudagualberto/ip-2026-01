package main

import "fmt"

func soma_reais(lista []float64) float64 {
	if len(lista) == 0 {
		return 0
	}
	return lista[0] + soma_reais(lista[1:])
}

func main() {
	lista_reais := []float64{1, 2, 3, 4, 5}
	s := soma_reais(lista_reais)
	fmt.Println("A soma é", s)
}
