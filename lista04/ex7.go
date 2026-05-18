package main

import "fmt"

func main() {
	vetor := []int{}
	// vetor = append(vetor, 1)
	for i:=1; i<=200; i+=2 {
		vetor = append(vetor, i)
	}
	fmt.Println(vetor)
}