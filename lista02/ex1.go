package main

import "fmt"

var n int

func main() {
	fmt.Println("Digite um valor inteiro:")
	fmt.Scan(&n)
	if n%2 == 0 {
		fmt.Println("O número é par")
	} else {
		fmt.Println("O número é ímpar")
	}
}
