package main

import (
	"fmt"
	"strconv"
)

var n1, n2, n3, conc, n string
var concInt, quadrado int

func main() {
	numero := make([]string, 3)

	for i := 0; i < 3; i++ {
		fmt.Println("Digite um número: ")
		fmt.Scan(&numero[i])
		if len(numero[i]) > 1 {
			fmt.Println("Número inválido")
			break
		}
	}
	conc = numero[0] + numero[1] + numero[2]
	concInt, _ = strconv.Atoi(conc)
	quadrado = concInt * concInt
	fmt.Printf("O número formado é %v e seu quadrado é %v", conc, quadrado)

}
