package main

import "fmt"

var x, y, z, w, det float64
var matriz [2][2]int

func main() {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println("Digite um valor")
			fmt.Scan(&matriz[i][j])
		}
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(matriz[i][j], " ")
		}
		fmt.Println()
	}
	det = (float64(matriz[0][0]) * float64(matriz[1][1])) - (float64(matriz[0][1]) * float64(matriz[1][0]))
	fmt.Println("DETERMINANTE: ", det)

}
