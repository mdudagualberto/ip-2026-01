package main

import "fmt"

func main() {
	vetor := []int{}
	num1F := 0
	num2F := 1
	for i:=0; i < 50; i++ {
		num1F, num2F = num1F + num2F, num1F
		vetor = append(vetor, num1F)
	}
	fmt.Println(vetor)
}