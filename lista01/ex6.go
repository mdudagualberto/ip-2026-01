package main

import "fmt"

var testes int
var f float64

func main() {
	fmt.Println("Quantas temperaturas você quer converter?")
	fmt.Scan(&testes)
	for i := 1; i <= testes; i++ {
		fmt.Println("\nDiga a temperatura em Fahrenheit: ")
		fmt.Scan(&f)
		c := 5 * (f - 32) / 9
		fmt.Printf("%v Fahrenheit equivale a %v Celsius", f, c)
	}
}
