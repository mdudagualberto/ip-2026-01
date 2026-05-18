package main

import "fmt"

var a, b, c, delta float64

func main() {
	fmt.Println("COEFICIENTES DA EQUAÇÃO")
	fmt.Println("Valor de A:")
	fmt.Scan(&a)
	fmt.Println("Valor de B:")
	fmt.Scan(&b)
	fmt.Println("Valor de C:")
	fmt.Scan(&c)
	delta = (b * b) - (4 * a * c)
	fmt.Printf("%.2f\n", delta)
}
