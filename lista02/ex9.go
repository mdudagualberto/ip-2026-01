package main

import "fmt"

var compra, venda float64

func main() {
	fmt.Println("Digite o valor da compra:")
	fmt.Scan(&compra)
	if compra < 10 {
		venda = compra * 1.7
	} else if compra >= 10 && compra < 30 {
		venda = compra * 1.5
	} else if compra >= 30 && compra < 50 {
		venda = compra * 1.4
	} else {
		venda = compra * 1.3
	}
	fmt.Printf("O valor de venda foi de %v reais\n", venda)
}
