package main

import "fmt"

var nome_aluno, conceito, situacao string
var nota1, nota2, nota3, media_ex, media_final float64

func main() {
	fmt.Println("Nome do aluno:")
	fmt.Scan(&nome_aluno)
	fmt.Println("Digite as três notas do aluno:")
	fmt.Scan(&nota1, &nota2, &nota3)
	fmt.Println("Digite a média dos exercícios do aluno:")
	fmt.Scan(&media_ex)
	media_final = (nota1 + (nota2 * 2) + (nota3 * 3) + media_ex) / 7
	if media_final >= 9 && media_final <= 10 {
		conceito = "A"
	} else if media_final < 9 && media_final >= 7.5 {
		conceito = "B"
	} else if media_final < 7.5 && media_final >= 6 {
		conceito = "C"
	} else if media_final < 6 && media_final >= 4 {
		conceito = "D"
	} else {
		conceito = "E"
	}
	if conceito == "A" || conceito == "B" || conceito == "C" {
		situacao = "APROVADO"
	} else {
		situacao = "REPROVADO"
	}
	fmt.Printf("A média do aluno %v foi de %v pontos, obtendo conceito %v e situação final de %v\n", nome_aluno, media_final, conceito, situacao)
}
