package main

import "fmt"

var salario, novoSalario float64

func main() {
	fmt.Println("Qual é seu salário?")
	fmt.Scan(&salario)
	if salario <= 300 {
		novoSalario = salario * 1.50
	} else {
		novoSalario = salario * 1.30
	}
	fmt.Printf("O NOVO SALÁRIO É: %.2f\n", novoSalario)
}
