package main

import "fmt"

var nprimo int

func primos(n int) bool {
	for i:=2; i<n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func main() {
	vetorPrimos := []int{}
	for i:=0; i < 10; i++ {
		fmt.Scan(&nprimo)
		vetorPrimos = append(vetorPrimos, nprimo)
	}
	for i:=0; i < 10; i++ {
		if primos(vetorPrimos[i]) {
			fmt.Printf("%v posição %v\n", vetorPrimos[i], i)
		}
	}
}