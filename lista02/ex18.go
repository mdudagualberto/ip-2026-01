package main

import "fmt"

var preco_normal_dvd float64
var dia_semana_aluguel, categoria_dvd string

func main() {
	fmt.Println("Qual é o preço normal do DVD?")
	fmt.Scan(&preco_normal_dvd)
	fmt.Println("Qual é a categoria do DVD?")
	fmt.Scan(&categoria_dvd)
	fmt.Println("Que dia da semana você irá comprar?")
	fmt.Scan(&dia_semana_aluguel)
	if categoria_dvd == "lançamento" {
		preco_normal_dvd *= 1.15
		if dia_semana_aluguel == "2" || dia_semana_aluguel == "3" || dia_semana_aluguel == "5" {
			preco_normal_dvd *= 0.6
		} else {
			preco_normal_dvd = preco_normal_dvd
		}
	} else if categoria_dvd == "comum" {
		if dia_semana_aluguel == "2" || dia_semana_aluguel == "3" || dia_semana_aluguel == "5" {
			preco_normal_dvd *= 0.6
		} else {
			preco_normal_dvd = preco_normal_dvd
		}

	}
	fmt.Printf("Você pagou %v reais no DVD de categoria %v\n", preco_normal_dvd, categoria_dvd)
}
