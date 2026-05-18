package main

import "fmt"

func inverte(array []int) []int {
	if len(array) == 0 {
		return []int{}
	}
	return append(inverte(array[1:]), array[0])
}

func main() {
	array := []int{1, 2, 3, 4, 5}
	i := inverte(array)
	fmt.Println(i)
}
