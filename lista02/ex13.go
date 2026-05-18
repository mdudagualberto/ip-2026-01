package main

import (
	"fmt"
	"strconv"
)

var nInteiro int

func main() {
	fmt.Println("Digite um número inteiro de 3 casas:")
	fmt.Scan(&nInteiro)
	if nInteiro < 100 || nInteiro > 999 {
		fmt.Println("Esse número não tem 3 casas")
	} else {
		str := strconv.Itoa(nInteiro)
		fmt.Println(str)
		letra := string(str[1])
		fmt.Println("O dígito das dezenas é", letra)

	}
}
