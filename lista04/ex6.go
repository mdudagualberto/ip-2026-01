package main

import "fmt"

func main() {
	vetor := []int{}

	for i:=100; i >=1; i-- {
		vetor = append(vetor, i)
	}

	fmt.Println(vetor)
}