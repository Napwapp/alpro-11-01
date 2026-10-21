package main

import "fmt"

func main() {
	var n, nilai int
	var nama string

	var lowest, secondLowest string
	min1 := 999999
	min2 := 999999

	fmt.Scan(&n)

	i := 1
	for i <= n {
		fmt.Print("Nama:")
		fmt.Scan(&nama)

		fmt.Print("Nilai:")
		fmt.Scan(&nilai)

		if nilai < min1 {
			min2 = min1
			secondLowest = lowest

			min1 = nilai
			lowest = nama
		} else if nilai < min2 {
			min2 = nilai
			secondLowest = nama
		}

		i++
	}

	fmt.Printf("\nSiswa dengan Nilai Terendah: %v (%v)\n", lowest, min1)
	fmt.Printf("Siswa dengan Nilai Terendah ke-2: %v (%v)\n", secondLowest, min2)
}