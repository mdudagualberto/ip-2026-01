// package main

// import "fmt"

// func incrementa(x int) {
// 	x++
// }

// func main() {
// 	v :=0
// 	incrementa(v)
// 	fmt.Println(v)
// }

package main

import "fmt"

func incrementaa(x *int) {
	*x++
}

func main() {
	v := 0
	incrementaa(&v)
	fmt.Println(v)
}
