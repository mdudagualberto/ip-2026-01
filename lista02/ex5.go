package main

import "fmt"

var num int

func main() {
	fmt.Println("Digite um valor inteiro:")
	fmt.Scan(&num)
	if num%5 == 0 {
		fmt.Println("O número é divisível por 5")
	} else {
		fmt.Println("O número não é divisível por 5")
	}
}
