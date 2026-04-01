package main

import "fmt" 

var numero int

func fatorial(numero int) int {
	if numero == 1 {
		return numero
	}
	return numero * fatorial(numero-1)
}

func main() {
	fmt.Println("Digite o número que você quer o fatorial:")
	fmt.Scan(&numero)
	f := fatorial(numero)
	fmt.Println("O fatorial é", f)
}
