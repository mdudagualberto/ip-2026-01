package main

import "fmt"

var f, polegadas float64

func main() {
	fmt.Println("\nDiga a temperatura em Fahrenheit: ")
	fmt.Scan(&f)
	c := 5 * (f - 32) / 9
	fmt.Println("\nDiga a quantidade de chuva em polegadas: ")
	fmt.Scan(&polegadas)
	mm := polegadas / 25.4
	fmt.Printf("%v Fahrenheit equivale a %v Celsius", f, c)
	fmt.Printf("\nQUANTIDADE DE CHUVA: %v mm", mm)

}
