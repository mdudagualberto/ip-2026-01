package main

import "fmt"

var matricula, qntd_extras int
var salario_extra, salario_bruto, desconto_inss, desconto_imposto, salario_liquido float64

func main() {
	salario_min := 788.00
	valor_extra := 10.00
	fmt.Println("Matrícula do funcinário:")
	fmt.Scan(&matricula)
	fmt.Println("Quantidade de horas extras que o funcionário trabalhou:")
	fmt.Scan(&qntd_extras)
	salario_extra = float64(qntd_extras) * valor_extra
	salario_bruto = 3*salario_min + salario_extra
	if salario_bruto > 1500 {
		desconto_inss = 0.12 * salario_bruto
	}
	if salario_bruto > 2000 {
		desconto_imposto = 0.2 * salario_bruto
	}
	salario_liquido = salario_bruto - desconto_imposto - desconto_inss
	fmt.Printf("O salário líquido do funcionário de matrícula %v é de %v reais\n", matricula, salario_liquido)
}
