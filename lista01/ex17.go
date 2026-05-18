package main

import "fmt"

var num3, num4 int

func main() {
	fmt.Println("Digite dois números:")
	fmt.Scan(&num3, &num4)
	if num3%2 == 0 {
		for i := 1; i <= num4; i++ {
			fmt.Println(num3)
			num3 += 2
		}
	} else {
		fmt.Println("O primeiro número não é par\n")
	}
}
