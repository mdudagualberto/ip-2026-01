package main

import "fmt"

var inteiro int
var somaa float64

func main() {
	fmt.Println("Digite um valor inteiro, positivo e maior que 1:")
	fmt.Scan(&inteiro)
	if inteiro <= 1 {
		fmt.Println("Número inválido")
	}
	for k := 1; k <= inteiro; k++ {
		somaa += 1 / float64(k)
	}
	fmt.Printf("A soma é: %.6f\n", somaa)
}
