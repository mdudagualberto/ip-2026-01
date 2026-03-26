package main

import "fmt"

var idade int
var classificação string

func main() {
	fmt.Println("Digite sua idade:")
	fmt.Scan(&idade)
	if idade < 0 {
		fmt.Println("Essa idade não é válida")
	} else {
		if idade >= 0 && idade <= 2 {
			classificação = "recém-nascido"
		} else if idade >= 3 && idade <= 11 {
			classificação = "criança"
		} else if idade >= 12 && idade <= 19 {
			classificação = "adolescente"
		} else if idade >= 20 && idade <= 55 {
			classificação = "adulto"
		} else {
			classificação = "idoso"
		}
		fmt.Println("Sua classificação de idade é de", classificação)
	}
}
