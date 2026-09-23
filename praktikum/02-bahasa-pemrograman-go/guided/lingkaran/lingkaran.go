package main

import "fmt"

func main() {
	phi := 3.14

	var (
		r, luas float64
	)

	// Input
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scanln(&r)

	// Operasi
	luas = phi * r * r

	// Output
	fmt.Println(luas)
}
