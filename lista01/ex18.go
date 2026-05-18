package main

import "fmt"

var ValorInicial, razao, elementos, somaPA, valorN int

func main() {
	fmt.Println("Digite o valor inicial, a razão e o número de elementos da progressão aritmética:")
	fmt.Scan(&ValorInicial, &razao, &elementos)
	valorN = ValorInicial + (elementos-1)*razao
	somaPA = (ValorInicial + valorN) * elementos / 2
	println("SOMA DA PA: ", somaPA)
}
