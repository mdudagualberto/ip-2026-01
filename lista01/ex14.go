package main

import (
	"fmt"
	"math"
)

var alt, aresta, areaBase, volume float64

func main() {
	fmt.Println("Digite a altura da pirâmide:")
	fmt.Scan(&alt)
	fmt.Println("Digite a aresta da base da pirâmide:")
	fmt.Scan(&aresta)
	areaBase = 3 * (aresta * aresta) * math.Sqrt(3) / 2
	volume = areaBase * alt / 3
	fmt.Printf("\nO volume da pirâmide é %v metros cúbicos", volume)
}
