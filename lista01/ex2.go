package main

import "fmt"

var teste, publicoTotal, popular, geral, arquibancada, cadeiras int
var pPopular, pGeral, pArquibancada, pCadeiras int
var rendaPopular, rendaGeral, rendaArquibancada, rendaCadeiras, rendaTotal int

func main() {
	fmt.Println("Quer analisar quantos jogos?")
	fmt.Scan(&teste)
	for i := 1; i <= teste; i++ {
		fmt.Println("\nQuantas pessoas foram ao jogo? ")
		fmt.Scan(&publicoTotal)
		fmt.Println("Qual a porcentagem de pessoas na área popular? ")
		fmt.Scan(&pPopular)
		fmt.Println("Qual a porcentagem de pessoas na área geral? ")
		fmt.Scan(&pGeral)
		fmt.Println("Qual a porcentagem de pessoas na área da arquibancada? ")
		fmt.Scan(&pArquibancada)
		fmt.Println("Qual a porcentagem de pessoas nas cadeiras? ")
		fmt.Scan(&pCadeiras)
		popular = (pPopular * publicoTotal) / 100
		geral = (pGeral * publicoTotal) / 100
		arquibancada = (pArquibancada * publicoTotal) / 100
		cadeiras = (pCadeiras * publicoTotal) / 100
		rendaPopular = popular * 1
		rendaGeral = geral * 5
		rendaArquibancada = arquibancada * 10
		rendaCadeiras = cadeiras * 20
		rendaTotal = rendaPopular + rendaArquibancada + rendaGeral + rendaCadeiras
		fmt.Printf("A renda do jogo número %v é %v reais", i, rendaTotal)

	}

}
