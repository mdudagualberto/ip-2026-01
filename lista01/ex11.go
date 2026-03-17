package main

import "fmt"

var n int

func main() {
	fmt.Println("Digite um número: ")
	fmt.Scan(&n)
	if n%3 == 0 && n%5 == 0 {
		fmt.Println("\nO número é divisível")
	} else {
		fmt.Println("\nO número não é divisível")
	}
}
