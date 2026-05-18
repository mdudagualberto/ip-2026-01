package main

import "fmt"

var n1, n2 int

func main() {
	vetor1 := []int{}
	vetor2 := []int{}

	for i:=0; i<10; i++ {
		fmt.Scan(&n1)
		vetor1 = append(vetor1, n1)
	}

	for i:=0; i<5; i++ {
		fmt.Scan(&n2)
		vetor2 = append(vetor2, n2)
	}

	vetorResultante1 := []int{}
	vetorResultante2 := []int{}

	for i:=0; i<10; i++ {
		if vetor1[i] % 2 == 0 {
			for j, _ := range(vetor2) {
				vetor1[i] += vetor2[j]
			}
			vetorResultante1 = append(vetorResultante1, vetor1[i])
		} else {
			for j, _ := range(vetor2) {
				vetor1[i] += vetor2[j]
			}
			vetorResultante2 = append(vetorResultante2, vetor1[i])
		}
	}
	fmt.Println(vetorResultante1)
	fmt.Println(vetorResultante2)
}