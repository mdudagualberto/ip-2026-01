package main

import "fmt"

var horas, totalPago int

func main() {
	fmt.Println("Quantas horas você usou a charrete?")
	fmt.Scan(&horas)
	if horas%3 == 0 {
		totalPago = (horas / 3) * 10
	} else if horas%3 != 0 {
		totalPago = (horas%3)*5 + (horas/3)*10
	}
	fmt.Println("VALOR A SER PAGO: ", totalPago)
}
