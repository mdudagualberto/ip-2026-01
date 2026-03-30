package main

import "fmt"

var valores, ssoma int

func main() {
	array := []int{}
	for i := 0; i < 5; i++ {
		fmt.Println("Digite um valor:")
		fmt.Scan(&valores)
		array = append(array, valores)
	}
	for i := 0; i < 5; i++ {
		ssoma += array[i]
	}
	fmt.Println("O array é:", array)
	fmt.Println("A soma entre os elementos do array é:", ssoma)
}
