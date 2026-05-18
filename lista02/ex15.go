package main

import "fmt"

var dia, mes, ano int
var mes_extenso string

func main() {
	fmt.Println("Digite o dia do seu nascimento:")
	fmt.Scan(&dia)
	fmt.Println("Digite o mês do seu nascimento:")
	fmt.Scan(&mes)
	fmt.Println("Digite o ano do seu nascimento:")
	fmt.Scan(&ano)
	meses := []string{"zero", "janeiro", "fevereiro", "março", "abril", "maio", "junho",
		"julho", "agosto", "setembro", "outubro", "novembro", "dezembro"}
	for indice := range len(meses) {
		if mes == indice {
			mes_extenso = meses[indice]
		}
	}
	fmt.Printf("Você nasceu no dia %v de %v de %v\n", dia, mes_extenso, ano)
}
