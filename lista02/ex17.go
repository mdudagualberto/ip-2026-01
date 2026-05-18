package main

import "fmt"

var tipoConsumidor, codigo string
var metrosCubicosGastos int
var totalGasto float64

func main() {
	fmt.Println("Qual é o código verificador da conta?")
	fmt.Scan(&codigo)
	fmt.Println("Você é consumidor residencial, comercial ou industrial?")
	fmt.Scan(&tipoConsumidor)
	fmt.Println("Quantos metros cúbicos você gastou?")
	fmt.Scan(&metrosCubicosGastos)
	if tipoConsumidor == "residencial" {
		totalGasto = 5 + (0.05 * float64(metrosCubicosGastos))
	} else if tipoConsumidor == "comercial" {
		if metrosCubicosGastos <= 80 {
			totalGasto = 500
		} else {
			totalGasto = 500 + (0.25 * (float64(metrosCubicosGastos - 500)))
		}
	} else if tipoConsumidor == "industrial" {
		if metrosCubicosGastos <= 100 {
			totalGasto = 800
		} else {
			totalGasto = 800 + (0.04 * (float64(metrosCubicosGastos - 100)))
		}
	}
	fmt.Println("CONTA: ", codigo)
	fmt.Println("VALOR DA CONTA: ", totalGasto)
}
