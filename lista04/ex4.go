package main

import "fmt"

var numero, vezes int

func contains(slice []int, value int) bool {
	for _, v := range(slice) {
		if v == value {
			return true
		}
	}
	return false
}

func contador(slice []int, value int) int {
	count := 0
	for _, v := range(slice) {
		if v == value {
			count++
		}
	}
	return count
}

func main() {
	vetoor := []int{}
	repetidos := []int{}
	for i:=0; i<10; i++ {
		fmt.Scan(&numero)
		vetoor = append(vetoor, numero)
	}
	for _, v := range(vetoor) {
		if contador(vetoor, v) > 1 && !contains(repetidos, v) {
			repetidos = append(repetidos, v)
		}
	}
	for _, v := range(repetidos) {
		fmt.Printf("Valor %v repetido %v vezes\n", v, contador(vetoor, v))
	}
}