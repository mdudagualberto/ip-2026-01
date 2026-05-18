package main

import "fmt"

var a, b int

func main() {
	fmt.Println("Digite dois valores inteiros:")
	fmt.Scan(&a, &b)
	if a%b == 0 {
		fmt.Println("O primeiro número é divisível pelo segundo")
	} else {
		fmt.Println("O primeiro número não é divisível pelo segundo")
	}
}
