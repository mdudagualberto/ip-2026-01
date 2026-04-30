package main

import "fmt"

var altura, somaAlturas int

func main() {
	vetor := []int{}
	for i:=0; i<10; i++ {
		fmt.Scan(&altura)
		somaAlturas += altura
		vetor = append(vetor, altura)
	}
	mediaAlturas := somaAlturas/10
	for i:=0; i<10; i++ {
		if vetor[i] > mediaAlturas {
			fmt.Println(vetor[i])
		}
	}
}