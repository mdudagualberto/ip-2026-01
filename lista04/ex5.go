package main

import "fmt"

var numeroo int

func main() {
	menor := 10000000
	indiceMenor := 0

	vetor := []int{}
	for i:=0; i<10; i++ {
		fmt.Scan(&numeroo)
		vetor = append(vetor, numeroo)
	}

	for i, v := range(vetor) {
		if v < menor {
			menor = v
			indiceMenor = i
		}
 	}
	fmt.Printf("O menor elemento do vetor é %v e sua posição dentro do vetor é %v\n", menor, indiceMenor)
}