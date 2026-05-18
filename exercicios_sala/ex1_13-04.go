package main

import "fmt"

func buscaSequencial(lista []int, chave int) int {
	for i := 0; i <= len(lista) - 1; i++ {
		if chave == lista[i] {
			return i
		} else {
			continue
		}
	}
	return -1
}

var numeroLista int

func main() {
	lista := []int{1, 2, 3, 4, 5}
	fmt.Println(lista)
	fmt.Println("Digite um número da lista:")
	fmt.Scan(&numeroLista)
	fmt.Printf("O índice desse número é %v\n", buscaSequencial(lista, numeroLista))
}
