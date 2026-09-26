package main

import "fmt"

func main() {
	var (
		nama string
		skorBahasaInggris, skorMatematika int
		rataRata int
	)

	// Membaca input dari pengguna
	fmt.Scanln(&nama)
	fmt.Scanln(&skorBahasaInggris)
	fmt.Scanln(&skorMatematika)

	// Operasi
	totalSkor := skorBahasaInggris + skorMatematika
	rataRata = (totalSkor) / 2

	// Output
	fmt.Println(nama)
	fmt.Println(totalSkor)
	fmt.Println(rataRata)
}
