package main

import "fmt"

var id int
var classe string

func main() {
	fmt.Println("Qual é sua idade?")
	fmt.Scan(&id)
	if id < 16 {
		classe = "não-eleitor"
	} else if id >= 18 && id <= 65 {
		classe = "eleitor obrigatório"
	} else if id > 65 || (id >= 16 && id < 18) {
		classe = "eleitor facultativo"
	}
	fmt.Println("Sua classificação eleitoral é de", classe)
}
