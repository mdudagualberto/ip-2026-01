package main

import "fmt"

var destino, ida_volta, valor_viagem int

func main() {
	fmt.Println("Os destinos são:\n1- Região Norte\n2- Região Nordeste\n3- Região Centro-Oeste\n4- Região Sul")
	fmt.Println("Digite um número de 1 a 4 para escolher seu destino:")
	fmt.Scan(&destino)
	if destino < 1 || destino > 4 {
		fmt.Println("Essa opção não é válida")
	} else {
		fmt.Println("Digite '1' se a viagem for de ida e volta e '2' se a viagem for só de ida:")
		fmt.Scan(&ida_volta)
		if ida_volta < 1 || ida_volta > 2 {
			fmt.Println("Essa opção não é válida")
		} else {
			if ida_volta == 2 {
				if destino == 1 {
					valor_viagem = 500
				} else if destino == 2 || destino == 3 {
					valor_viagem = 350
				} else {
					valor_viagem = 300
				}

			} else {
				if destino == 1 {
					valor_viagem = 900
				} else if destino == 2 {
					valor_viagem = 650
				} else if destino == 3 {
					valor_viagem = 600
				} else {
					valor_viagem = 550
				}

			}
		}
	}
	fmt.Printf("O valor da viagem foi de %v reais\n", valor_viagem)
}
