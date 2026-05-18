package main

import "fmt"

var num1, num2, num3 int

func maior(num1, num2, num3 int) int {
	if num1 > num2 {
		num1, num2 = num2, num1
	} 
	if num1 > num3 {
		num1, num3 = num3, num1
	} 
	if num2 > num3 {
		num2, num3 = num3, num2
	}
	return num3
}

func main() {
	fmt.Println("Digite o primeiro número:")
	fmt.Scan(&num1)
	fmt.Println("Digite o segundo número:")
	fmt.Scan(&num2)
	fmt.Println("Digite o terceiro número:")
	fmt.Scan(&num3)
	maior := maior(num1, num2, num3)
	fmt.Println("O maior número é", maior)
}
