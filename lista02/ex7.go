package main

import "fmt"

var num1, num2, num3 int

func main() {
	fmt.Println("Digite três números:")
	fmt.Scan(&num1, &num2, &num3)
	if num1 > num2 {
		num1, num2 = num2, num1
	} else if num1 > num3 {
		num1, num3 = num3, num1
	} else if num2 > num3 {
		num2, num3 = num3, num2
	}
	menor := num1
	inter := num2
	maior := num3
	fmt.Println("Os valores em ordem crescente são:", menor, inter, maior)

}
