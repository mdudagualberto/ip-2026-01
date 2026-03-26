package main

import "fmt"

var etiqueta, parcelas float64
var condicao_pagamento string
var qntd_parcelas int

func main() {
	fmt.Println("Qual é o valor da etiqueta?")
	fmt.Scan(&etiqueta)
	fmt.Println("OPÇÕES DE PAGAMENTO\n1)À vista, dinheiro ou cheque, 10% de desconto\n2)À vista, cartão de crédito, 5% de desconto\n3)Em 2 vezes, preço normal de etiqueta sem juros\n4)Em 3 vezes, preço normal de etiqueta + 10% de juros")
	fmt.Println("Selecione uma opção de pagamento:")
	fmt.Scan(&condicao_pagamento)
	if condicao_pagamento == "1" {
		parcelas = etiqueta * 0.9
		qntd_parcelas = 1
	} else if condicao_pagamento == "2" {
		parcelas = etiqueta * 0.95
		qntd_parcelas = 1
	} else if condicao_pagamento == "3" {
		parcelas = etiqueta / 2
		qntd_parcelas = 2
	} else {
		parcelas = (etiqueta / 3) * 1.1
		qntd_parcelas = 3
	}
	fmt.Printf("Você irá pagar %v parcelas de %v reais\n", qntd_parcelas, parcelas)
}
