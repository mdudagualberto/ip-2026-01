package main

import "fmt"

var num, quad int

func main() {
	fmt.Println("Digite um valor entre 5 e 2000")
	fmt.Scan(&num)
	if num < 5 || num > 2000 {
		fmt.Println("Número inválido")
	}
	for i := 6; i < num; i++ {
		if i%2 == 0 {
			quad = i * i
			fmt.Printf("\n%v^2: %v", i, quad)
		}
	}
}
