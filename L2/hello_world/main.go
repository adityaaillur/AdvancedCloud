package main

import (
	"fmt"
)

const PI = 3.14
const C = 1

func main() {
	person1 := person{"aditya ", "illur"}
	fmt.Println("Person's full name:", person1.fullName())

	fb()

	// fmt.Println(quote.Go())
	// power := 9000
	// fmt.Printf("its over %d", power)
	// fmt.Println()
	// fmt.Println(quote.Hello())
	// fmt.Println(add(1, 2))
	// fmt.Printf("its over %d", power)
}

// func add(a, b int) (c int) {
// 	c = a + b
// 	return c
// }

// func num(a, b int) (c int, d int) {
// 	return a, b
// }
