package main

import "fmt"

func main() {
	var (
		jumlahTiket, jumlahTransaksi, totalPendapatan, harga, diskon int
		jenis, hari                                                  string
	)

	fmt.Scan(&jumlahTransaksi)
	for i := 1; i <= jumlahTransaksi; i++ {
		harga = 0
		fmt.Printf("Transaksi ke-%d\n", i)
		fmt.Print("Masukkan jenis film (Reguler, 3D, IMAX): ")
		fmt.Scan(&jenis)

		fmt.Print("Hari: ")
		fmt.Scan(&hari)

		fmt.Print("Jumlah Tiket: ")
		fmt.Scan(&jumlahTiket)

		if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
			if jenis == "Reguler" {
				harga = 40000
			} else if jenis == "3D" {
				harga = 60000
			} else if jenis == "IMAX" {
				harga = 80000
			} else {
				fmt.Println("Jenis film tidak valid")
			}
		} else if hari == "Jumat" || hari == "Sabtu" || hari == "Minggu" {
			if jenis == "Reguler" {
				harga = 50000
			} else if jenis == "3D" {
				harga = 70000
			} else if jenis == "IMAX" {
				harga = 90000
			} else {
				fmt.Println("Jenis film tidak valid")
			}
		} else {
			fmt.Println("Hari tidak valid")
		}

		totalHarga := harga * jumlahTiket

		// Diskon ketika beli tiket 5 keatas
		if jumlahTiket >= 5 {
			diskon = totalHarga * 10 / 100
			totalHarga = totalHarga - diskon
		}

		totalPendapatan = totalPendapatan + totalHarga

		fmt.Printf("\n====================\n")
		fmt.Println("Total Harga:",totalHarga)
	}

	fmt.Println("Total Pendapatan:",totalPendapatan)
	fmt.Println("====================")

}
