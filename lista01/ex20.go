package main

import "fmt"

var hora, min, seg, tot int

func main() {
	fmt.Println("Valor em horas:")
	fmt.Scan(&hora)
	fmt.Println("Valor em minutos:")
	fmt.Scan(&min)
	fmt.Println("Valor em segundos:")
	fmt.Scan(&seg)
	tot = (hora * 3600) + (min * 60) + seg
	fmt.Println("\nO valor em segundos é: ", tot)
}
