package main

import "fmt"

var numeroo int

func main() {
	fmt.Println("Digite um valor inteiro:")
	fmt.Scan(&numeroo)
	if numeroo > 20 && numeroo < 90 {
		fmt.Println("O valor está entre 20 e 90")
	} else {
		fmt.Println("O valor não está entre 20 e 90")
	}
}
