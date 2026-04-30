package main

import "fmt"

var n, maior50 int

func main() {
	vetor := []int{}
	for i:=0; i<10; i++ {
		fmt.Scan(&n)
		vetor = append(vetor, n)
	}
	maior50 = 0
	for i:=0; i<=9; i++ {
		if vetor[i] > 50 {
			fmt.Printf("Número %v, posição %v\n", vetor[i], i)
			maior50++
		}
	}
	if maior50 == 0 {
		fmt.Println("Não possui números maiores que 50")
	}
}