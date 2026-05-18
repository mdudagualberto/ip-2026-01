package main

import "fmt"

var num, somaPares, impares int

func main() {
	vetorr := []int{}
	vetorPares := []int{}
	vetorImpares := []int{}
	for i:=0; i<10; i++ {
		fmt.Scan(&num)
		vetorr = append(vetorr, num)
	}
	for _, v := range(vetorr) {
		if v % 2 == 0 {
			vetorPares = append(vetorPares, v)
			somaPares += v
		} else {
			vetorImpares = append(vetorImpares, v)
			impares++
		}
	}
	fmt.Println(vetorPares)
	fmt.Println(somaPares)
	fmt.Println(vetorImpares)
	fmt.Println(impares)
}