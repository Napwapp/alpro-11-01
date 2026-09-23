package main

import "fmt"

func main () {
	var (
		uangRupiah int
		lembar10rb, lembar5rb, lembar1rb int
	)

	// Input
	fmt.Println("Masukkan jumlah uang dalam rupiah: ")
	fmt.Scanln(&uangRupiah)

	// Operasi
	lembar10rb = uangRupiah / 10000 // bagi 10rb
	uangRupiah %= 10000 // sisa setelah dibagi 10rb
	lembar5rb = uangRupiah / 5000 // sisa tadi dibagi 5rb
	uangRupiah %= 5000 // sisa setelah dibagi 5rb
	lembar1rb = uangRupiah / 1000 // sisa tadi dibagi 1rb

	// Output
	fmt.Printf("\n%d Lembar uang 10000\n", lembar10rb)
	fmt.Printf("\n%d Lembar uang 5000\n", lembar5rb)
	fmt.Printf("\n%d Lembar uang 1000\n", lembar1rb)
}
