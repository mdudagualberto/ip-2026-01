package main

import "fmt"

var raio, altura, area, custo float64

func main() {
	fmt.Println("Qual o raio da lata em metros?")
	fmt.Scan(&raio)
	fmt.Println("Qual a altura da lata em metros?")
	fmt.Scan(&altura)
	area = 2*(3.14159*(raio*raio)) + (2 * 3.14159 * raio * altura)
	custo = 100 * area
	fmt.Println("VALOR DO CUSTO: ", custo)
}
