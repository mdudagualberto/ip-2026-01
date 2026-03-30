package main

import "fmt"

var valoress float64

func main() {
	lista := []float64{}
	for i := 0; i < 10; i++ {
		fmt.Println("Digite um valor:")
		fmt.Scan(&valoress)
		lista = append(lista, valoress)
	}
	novaLista := []float64{}
	for i := 9; i >= 0; i-- {
		novaLista = append(novaLista, lista[i])
	}
	fmt.Println("Os números em ordem decrescente são:", novaLista)
}
