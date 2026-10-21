package main

import "fmt"

func main() {
	var (
		n int
		biayaTahunIni, p6 float64
	)

	biayaTahunIni = 10000000
	i := 2
	fmt.Scan(&n)

	fmt.Printf("Tahun ke 1: %0.f\n", biayaTahunIni)
	for i <= n {
		p6 = biayaTahunIni * 0.06
		biayaTahunIni = biayaTahunIni + p6
		fmt.Printf("Tahun ke %d: %0.f\n", i, biayaTahunIni)
		i++
	}

	fmt.Printf("Biaya tahun ke %d: %0.f",n, biayaTahunIni)
}
