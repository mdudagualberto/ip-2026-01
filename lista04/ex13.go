package main

import "fmt"

var nEmpregado, mesesEmpregado, mesMaisRecente int

// func bubbleSort(array []int) []int {
// 	n := len(array)
// 	for i:=0; i < n - 1; i++ {
// 		for j:=0; j < n - i - 1; j++ {
// 			if array[j] > array[j+1] {
// 				array[j], array[j+1] = array[j+1], array[j]
// 			}
// 		}
// 	}
// 	return array
// }

func main() {
	mesMaisRecente = 100000
	meses := []int{}
	numeroEmpregado := []int{}
	for i:=0; i < 100 ; i++ {
		fmt.Scan(&nEmpregado)
		fmt.Scan(&mesesEmpregado)
		if nEmpregado == 0 && mesesEmpregado == 0 {
			break
		}
		numeroEmpregado = append(numeroEmpregado, nEmpregado)
		meses = append(meses, mesesEmpregado)
	}
	for i:=0; i < 3 && i < len(meses); i++ {
		indiceMenor := 0
		for k:=1; k < len(meses); k++ {
			if meses[k] < meses[indiceMenor] {
				indiceMenor = k
			}
		}
		fmt.Println(numeroEmpregado[indiceMenor])
		meses[indiceMenor] = 999999
	} 
}