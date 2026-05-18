package main

import "fmt"

var xis int

func main() {
	fmt.Println("Digite um valor inteiro:")
	fmt.Scan(&xis)
	res := 8 / (2 - xis)
	fmt.Println("O resultado é", res)
}
