package main

import "fmt"

var nume int

func main() {
	vetor := []int{}

	for i:=0; i<15; i++ {
		fmt.Scan(&nume)
		if nume < 0 {
			vetor = append(vetor, -1)
		} else {
			vetor = append(vetor, (nume *nume))
		}
	}
	fmt.Println(vetor)
}