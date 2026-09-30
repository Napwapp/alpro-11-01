package main

import "fmt"

func main() {
	var (
		kue, anggotaKeluarga, sisaKue int
	)

	fmt.Println("Masukkan kue yang dimasak:")
	fmt.Scanln(&kue)

	fmt.Println("\nMasukkan jumlah anggota keluarga:")
	fmt.Scanln(&anggotaKeluarga)

	fmt.Println("\nSisa kue:")
	sisaKue = kue % anggotaKeluarga 

	fmt.Println(sisaKue)
}
