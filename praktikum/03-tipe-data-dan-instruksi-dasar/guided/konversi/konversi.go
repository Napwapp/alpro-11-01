package main

import "fmt"

func main() {
	var (
		c, k float64
	)

	// Input
	fmt.Println("Masukkan suhu celcius")
	fmt.Scanln(&c)

	// Konversi suhu celcius ke kelvin
	k = c + 273

	// Output
	fmt.Print(k)
}
