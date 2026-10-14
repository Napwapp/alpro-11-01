package main

import "fmt"

func main() {
	var (
		jumlahTiket, jumlahTransaksi, totalPendapatan, harga, diskon int
		jenis, hari                                                  string
	)

	fmt.Scan(&jumlahTransaksi)
	for i := 1; i <= jumlahTransaksi; i++ {
		fmt.Printf("Transaksi ke-%d\n", i)
		fmt.Println("Masukkan jenis film (Reguler, 3D, IMAX): ")
		fmt.Scan(&jenis)

		fmt.Println("Hari: ")
		fmt.Scan(&hari)

		fmt.Println("Jumlah Tiket")
		fmt.Scan(&jumlahTiket)

		if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
			if jenis == "Reguler" {
				harga = 40000
			} else if jenis == "3D" {
				harga = 60000
			} else if jenis == "IMAX" {
				harga = 80000
			}
		}

		if hari == "Jumat" || hari == "Sabtu" || hari == "Minggu" {
			if jenis == "Reguler" {
				harga = 50000
			} else if jenis == "3D" {
				harga = 70000
			} else if jenis == "IMAX" {
				harga = 90000
			}
		}

		totalHarga := harga * jumlahTiket

		// Diskon ketika beli tiket 5 keatas
		if jumlahTiket >= 5 {
			diskon = totalHarga * 10 / 100
			totalHarga = totalHarga - diskon
		}

		totalPendapatan = totalPendapatan + totalHarga

		fmt.Printf("Total Harga: %d\n", totalHarga)
		fmt.Printf("Total Pendapatan: %d\n", totalPendapatan)
	}

}
