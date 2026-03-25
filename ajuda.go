package main

import "fmt"

var nnotas int
var nota, soma, media float64

func main() {
	fmt.Println("Quantas notas quer digitar?")
	fmt.Scan(&nnotas)
	slice := []float64{}
	for i := 1; i <= nnotas; i++ {
		fmt.Println("Digite sua nota:")
		fmt.Scan(&nota)
		slice = append(slice, nota)
	}
	for i := 0; i < nnotas; i++ {
		soma += slice[i]
	}
	media = soma / float64(len(slice))
	fmt.Println(slice)
	fmt.Println(soma)
	fmt.Println(media)
}
