package main

import "fmt"

func main() {
	var tanggal, bulan, tahun int

	fmt.Print("Masukkan Tanggal: ")
	fmt.Scan(&tanggal)

	fmt.Print("Masukkan Bulan: ")
	fmt.Scan(&bulan)

	fmt.Print("Masukkan Tahun: ")
	fmt.Scan(&tahun)

	// Tahun kabisat
	kabisat := (tahun%4 == 0 && tahun%100 != 0) || tahun%400 == 0

	// Hari maksimum pada bulan tersebut
	hariMaks := 0
	
	switch bulan {
	case 1, 3, 5, 7, 8, 10, 12:
		hariMaks = 31
	case 4, 6, 9, 11:
		hariMaks = 30
	case 2:
		if kabisat {
			hariMaks = 29
		} else {
			hariMaks = 28
		}
	}

	// Validasi tanggal
	if tahun <= 0 || hariMaks == 0 || tanggal < 1 || tanggal > hariMaks {
		fmt.Println("Tidak valid")
		return
	}

	// Hitung tanggal besok
	if tanggal < hariMaks {
		tanggal++
	} else {
		tanggal = 1
		if bulan == 12 {
			bulan = 1
			tahun++
		} else {
			bulan++
		}
	}

	fmt.Println("Valid")
	fmt.Printf("%d/%d/%d\n", tanggal, bulan, tahun)
}