package main

import "fmt"

var nota_1, nota_2, nota3 int
var mediaa, somaa float64

func main() {
	slice := []float64{}
	fmt.Println("Digite a primeira nota:")
	fmt.Scan(&nota_1)
	slice = append(slice, float64(nota_1))
	fmt.Println("Digite a segunda nota:")
	fmt.Scan(&nota_2)
	slice = append(slice, float64(nota_2))
	fmt.Println("Digite a terceira nota:")
	fmt.Scan(&nota3)
	slice = append(slice, float64(nota3))
	for i := 0; i < 3; i++ {
		somaa += slice[i]
		mediaa = somaa / 3
	}
	fmt.Printf("A média das três notas é %v\n", mediaa)
	if nota_1 >= int(mediaa) {
		fmt.Println("A primeira nota está acima da média")
	} else {
		fmt.Println("A primeira nota não está acima da média")
	}
	if nota_2 >= int(mediaa) {
		fmt.Println("A segunda nota está acima da média")
	} else {
		fmt.Println("A segunda nota não está acima da média")
	}
	if nota3 >= int(mediaa) {
		fmt.Println("A terceira nota está acima da média")
	} else {
		fmt.Println("A terceira nota não está acima da média")
	}
}
