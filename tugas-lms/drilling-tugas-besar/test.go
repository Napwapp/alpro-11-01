package main

import "fmt"

func main() {
	var (
		kuota, jumlah, harga, total, terjual, pendapatan, transaksi, kategori  int
		hari, jenisHari, namaKategori string
	)

	fmt.Print("Masukkan kuota tiket: ")
	fmt.Scan(&kuota)
	
	terjual = 0
	pendapatan = 0
	transaksi = 0

	for kuota > 0 {
		fmt.Println("Sisa kuota:", kuota)

		for {
			fmt.Print("Kategori tiket untuk (1 = Anak, 2 = Dewasa): ")
			fmt.Scan(&kategori)

			if kategori == 1 || kategori == 2 {
				break
			}
			fmt.Println("kategori tidak valid (Masukkan angka 1 atau 2)")
		}

		for {
			fmt.Print("Masukkan hari (Senin - Minggu): ")
			fmt.Scan(&hari)

			if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" || hari == "Jumat" || hari == "Sabtu" || hari == "Minggu" {
				break
			}
			fmt.Println("Hari tidak valid (Senin, Selasa, Rabu, Kamis, Jumat, Sabtu, Minggu)")
		}

		for {
			fmt.Print("Masukkan jumlah tiket: ")
			fmt.Scan(&jumlah)

			if jumlah <= 0 {
				fmt.Println("jumlah tiket harus lebih dari 0")
				continue
			}
			if jumlah > kuota {
				fmt.Println("kuota tidak mencukupi")
				continue
			}
			break
		}

		switch hari {
		case "Senin", "Selasa", "Rabu", "Kamis", "Jumat":
			jenisHari = "hari biasa"
		case "Sabtu", "Minggu":
			jenisHari = "akhir pekan"
		}

		switch kategori {
		case 1:
			namaKategori = "anak"
			if jenisHari == "hari biasa" {
				harga = 20000
			} else {
				harga = 30000
			}
		case 2:
			namaKategori = "dewasa"
			if jenisHari == "hari biasa" {
				harga = 35000
			} else {
				harga = 50000
			}
		}

		total = harga * jumlah
		kuota = kuota - jumlah
		terjual = terjual + jumlah
		pendapatan = pendapatan + total
		transaksi = transaksi + 1

		fmt.Println("Bukti pembelian ke-", transaksi)
		fmt.Println("Kategori:", namaKategori)
		fmt.Println("Hari:", hari)
		fmt.Println("Jumlah tiket:", jumlah)
		fmt.Println("Harga satuan:", harga)
		fmt.Println("Total bayar:", total)
		fmt.Println("Sisa kuota:", kuota)
		fmt.Println()
	}

	fmt.Println("Tiket habis terjual")
	fmt.Println("Sisa kuota tiket:", kuota)
	fmt.Println("Jumlah tiket terjual:", terjual)
	fmt.Println("Total pendapatan:", pendapatan)
}