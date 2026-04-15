package main

import "fmt"

var base, expoentee int

func elevar(base, expoente int) int {
	if expoente == 0 {
		return 1
	}
	return base * elevar(base, (expoente-1))
}

func main() {
	fmt.Println("Digite a base:")
	fmt.Scan(&base)
	fmt.Println("Digite o expoente:")
	fmt.Scan(&expoentee)
	r := elevar(base, expoentee)
	fmt.Println("O resultado da exponenciação é", r)
}