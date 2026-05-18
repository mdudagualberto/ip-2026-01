package main

import "fmt"

var nDecimal int

func transformaBinario(n int) string {
	if n == 0 {
		return ""
	}
	return transformaBinario(n/2) + fmt.Sprint(n%2)
}

func main() {
	fmt.Println("Digite o número em decimal:")
	fmt.Scan(&nDecimal)
	b := transformaBinario(nDecimal)
	fmt.Println(b)
}