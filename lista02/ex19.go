package main

import (
	"fmt"
	"math"
)

var volume_cone, area_cone, volume_cilindro, area_cilindro, volume_esfera, area_esfera float64
var raio, altura float64
var opcao string

func main() {
	fmt.Println("Você quer calcular o volume e a área da superfície de: \n1) um cone reto\n2) um cilindro\n3) uma esfera")
	fmt.Scan(&opcao)
	fmt.Println("Digite o raio:")
	fmt.Scan(&raio)
	fmt.Println("Digite a altura:")
	fmt.Scan(&altura)
	volume_cone = (3 * (raio * raio) * altura) / 3
	area_cone = 3 * raio * math.Sqrt((raio*raio)+(altura*altura))
	volume_cilindro = 3 * (raio * raio) * altura
	area_cilindro = 2 * 3 * raio * altura
	volume_esfera = (4 * 3 * (raio * raio * raio)) / 3
	area_esfera = 4 * 3 * (raio * raio)
	if opcao == "1" {
		fmt.Printf("O volume é %v e a área é %v\n", volume_cone, area_cone)
	} else if opcao == "2" {
		fmt.Printf("O volume é %v e a área é %v\n", volume_cilindro, area_cilindro)
	} else {
		fmt.Printf("O volume é %v e a área é %v\n", volume_esfera, area_esfera)
	}
}
