package main

import "fmt"

var nota1, nota2, nota3, media float64

func main() {
	fmt.Println("Primeira nota: ")
	fmt.Scan(&nota1)
	fmt.Println("Segunda nota: ")
	fmt.Scan(&nota2)
	fmt.Println("Terceira nota: ")
	fmt.Scan(&nota3)
	fmt.Printf("\nAs notas são: %v, %v e %v", nota1, nota2, nota3)
	media = (nota1 + nota2 + nota3) / 3
	fmt.Printf("\nMédia: %v", media)
	if media >= 6 {
		fmt.Println("\nAPROVADO")
	} else {
		fmt.Println("\nREPROVADO")
	}
}
