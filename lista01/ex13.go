package main

import "fmt"

var nota float64
var parametro string

func main() {
	fmt.Println("Digite sua nota:")
	fmt.Scan(&nota)
	if nota >= 9 && nota <= 10 {
		parametro = "A"
	} else if nota < 9 && nota >= 7.5 {
		parametro = "B"
	} else if nota < 7.5 && nota >= 6 {
		parametro = "C"
	} else {
		parametro = "D"
	}
	fmt.Printf("\nNOTA: %v, CONCEITO: %v\n", nota, parametro)
}
