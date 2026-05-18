package main

import "fmt"

var x, y, z float64

func media(x, y, z float64) float64 {
	return (x + y + z)/3
}

func main() {
	fmt.Println("Digite o primeiro número:")
	fmt.Scan(&x)
	fmt.Println("Digite o segundo número:")
	fmt.Scan(&y)
	fmt.Println("Digite o terceiro número:")
	fmt.Scan(&z)
	m := media(x, y, z)
	fmt.Println("A média dos números é", m)
}
