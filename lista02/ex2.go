package main

import "fmt"

var numero int

func main() {
	fmt.Println("Digite um valor inteiro:")
	fmt.Scan(&numero)
	if numero > 0 {
		fmt.Println("O número é positivo")
	} else if numero < 0 {
		fmt.Println("O número é negativo")
	} else {
		fmt.Println("O número é nulo")
	}
}
