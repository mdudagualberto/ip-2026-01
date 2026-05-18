package main

import "fmt"

var n1, n2 int

func main() {
	fmt.Println("Digite dois números:")
	fmt.Scan(&n1, &n2)
	soma := n1 + n2
	if soma > 20 {
		soma += 8
		fmt.Println(soma)
	} else {
		soma -= 5
		fmt.Println(soma)
	}
}
