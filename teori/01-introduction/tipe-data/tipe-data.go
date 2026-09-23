package main

import "fmt"

func main() {
	// Number
	var (
		umur   int8
		anakKe int8
		saldo  int64
	)

	umur = 18
	anakKe = 2
	saldo = 1000000000000

	fmt.Println("Umur saya adalah", &umur, "tahun")
	fmt.Println("Anak ke-", anakKe)
	fmt.Println("Saldo saya", saldo)
	//  & untuk menampilkan alamat memori dari sebuah variabel

	// String
	var (
		namaDepan string
		namaBelakang string
	)

	namaDepan = "M Nawaf"
	namaBelakang = "Abduh"
	fmt.Println("Nama saya adalah", namaDepan, namaBelakang)
}
