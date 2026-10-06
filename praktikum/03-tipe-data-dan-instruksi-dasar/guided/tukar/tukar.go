package main

import "fmt"

func main() {
	var (x,y,z int)

	// Input nilai a dan b
	fmt.Print("Masukkan nilai x: ")
	fmt.Scanln(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scanln(&y)
	fmt.Print("Masukkan nilai z: ")
	fmt.Scanln(&z)

	// Tukar nilai x,y,z
	temp := x
	x = z
	z = y
	y = temp

	// Output hasil setelah ditukar
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}