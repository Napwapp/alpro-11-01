package main

import "fmt"

func main() {
	var (
		a int
		b int
	)

	// Input nilai a dan b
	fmt.Print("Masukkan nilai a: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scanln(&b)

	// Tukar nilai a dan b
	a, b = b, a

	// Output hasil setelah ditukar
	fmt.Println(a)
	fmt.Println(b)
}