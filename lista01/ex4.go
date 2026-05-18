package main

import "fmt"

var salarioMinimo, qntdKw int
var valorKw, valorPago, comDesconto float64

func main() {
	fmt.Println("Qual o valor do salário mínimo?")
	fmt.Scan(&salarioMinimo)
	fmt.Println("Qual foi a quantidade de Kw gasta por esta residência?")
	fmt.Scan(&qntdKw)
	valorKw = 0.007 * float64(salarioMinimo)
	valorPago = valorKw * float64(qntdKw)
	comDesconto = valorPago * 0.9
	fmt.Printf("\nCusto por Kw: R$ %v", valorKw)
	fmt.Printf("\nCusto do consumo: R$ %v", valorPago)
	fmt.Printf("\nCusto com desconto: R$%v", comDesconto)
}
