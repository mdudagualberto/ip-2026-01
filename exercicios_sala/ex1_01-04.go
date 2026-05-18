package main

import "fmt"

var n1, n2 float64

func soma(n1, n2 float64) float64 {
		return n1 + n2
	}

func main() {
	fmt.Println("Digite um valor:")
	fmt.Scan(&n1)
	fmt.Println("Digite outro número:")
	fmt.Scan(&n2)
	s := soma(n1, n2)
	fmt.Println("A soma dos números é", s)
}