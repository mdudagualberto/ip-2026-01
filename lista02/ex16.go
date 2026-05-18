package main

import (
	"fmt"
	"math"
)

var coef_a, coef_b, coef_c float64
var delta, raiz1, raiz2 float64

func main() {
	fmt.Println("Digite os coeficientes de uma equação de segundo grau:")
	fmt.Scan(&coef_a, &coef_b, &coef_c)
	delta = (float64(coef_b) * float64(coef_b)) - (4 * coef_a * coef_c)
	fmt.Println(delta)
	if delta < 0 {
		fmt.Println("Essa equação não tem raízes reais")
	} else {
		raiz1 = (-coef_b + math.Sqrt(delta)) / 2 * coef_a
		raiz2 = (-coef_b - math.Sqrt(delta)) / 2 * coef_a
		fmt.Printf("As raízes são %v e %v\n", raiz1, raiz2)
		if raiz1 == raiz2 {
			fmt.Println("RAÍZES IGUAIS")
		} else {
			fmt.Println("RAÍZES DIFERENTES")
		}
	}

}
