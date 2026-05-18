package main

import (
	"fmt"
	"math"
)

var x float64

func main() {
	fmt.Println("Digite um número:")
	fmt.Scan(&x)
	if x >= 0 {
		raiz := math.Sqrt(x)
		fmt.Println("A raiz quadrada do número é", raiz)
	} else {
		quadrado := x * x
		fmt.Println("O quadrado do número é", quadrado)
	}
}
