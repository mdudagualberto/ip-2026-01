package main

import "fmt"

type Pessoa struct {
	nome string
	altura float64
	pesoIdeal float64
}

var pessoa Pessoa
var nome string
var altura, pesoIdeal float64

func main() {
	lista := []Pessoa{}
	for {
		fmt.Println("Qual é seu nome?")
		fmt.Scan(&nome)
		if nome == "FIM" || nome == "fim" {
			break
		}
		fmt.Println("Qual sua altura?")
		fmt.Scan(&altura)
		pessoa.nome = nome
		pessoa.altura = altura
		pesoIdeal = (72.7 * pessoa.altura) - 58.0
		pessoa.pesoIdeal = pesoIdeal
		lista = append(lista, pessoa)
	}
	for i,valor := range(lista) {
		fmt.Printf("\nPessoa %v \n\tNome: %v \n\tAltura: %.2f \n\tPeso Ideal: %.2f \n", i+1, valor.nome, valor.altura, valor.pesoIdeal)
	}
	
}
