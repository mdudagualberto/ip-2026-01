package main

import (
	"fmt"
	"slices"
)

var valor_inicial float64
var escolha string
var nopcoes int

func main() {
	acessorios := []string{}
	fmt.Println("Qual é o valor inicial do carro?")
	fmt.Scan(&valor_inicial)
	fmt.Println("O carro pode ter as seguintes opções:\na) Ar condicionado (R$ 1750,00)\nb) Pintura metálica (R$ 800,00)\nc) Vidro elétrico (R$ 1200,00)\nd) Direção hidráulica (R$ 2000,00)")
	fmt.Println("Quantas opções você quer adicionar no carro?")
	fmt.Scan(&nopcoes)
	for i := 1; i <= nopcoes; i++ {
		fmt.Println("Digite a opção:")
		fmt.Scan(&escolha)
		acessorios = append(acessorios, escolha)
	}
	if slices.Contains(acessorios, "a") {
		valor_inicial += 1750
	}
	if slices.Contains(acessorios, "b") {
		valor_inicial += 800
	}
	if slices.Contains(acessorios, "c") {
		valor_inicial += 1200
	}
	if slices.Contains(acessorios, "d") {
		valor_inicial += 2000
	}
	fmt.Println(acessorios)
	fmt.Printf("O valor final é de %v reais", valor_inicial)
}
