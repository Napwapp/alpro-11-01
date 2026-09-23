package main

import "fmt"

func main () {
	a := 10
	b := 10

	a += b
	fmt.Println(a)

	// Unary operator
	i := 1
	i++
	fmt.Println(i)
	i++
	fmt.Println(i)
	i--
	fmt.Println(i)
	i--
	fmt.Println(i)
}
