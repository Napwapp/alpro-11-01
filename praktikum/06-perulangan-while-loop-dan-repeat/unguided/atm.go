package main

import "fmt"

func main() {
	var (
		password string
		pilihan,jumlah int
	)
	const passwordBenar = "1234"

	saldo := 10000000
	i := 0
	loginBerhasil := false

	// Perulangan login: berhenti jika password benar atau i habis
	for !loginBerhasil && i < 3 {
		fmt.Print("Masukkan password ATM: ")
		fmt.Scan(&password)

		if password == passwordBenar {
			loginBerhasil = true
		} else {
			i++
			fmt.Printf("\nPassword salah. Silakan coba lagi.\n")
		}
	}

	// Jika 3 kali salah, ATM terblokir
	if !loginBerhasil {
		fmt.Println("ATM terblokir. Program selesai.")
		return
	}

	fmt.Printf("\nLogin berhasil!\n")

	for pilihan != 3 {
		fmt.Println("\n==================")
		fmt.Println("Menu ATM:")
		fmt.Println("1. Lihat Saldo")
		fmt.Println("2. Tarik Uang")
		fmt.Println("3. Keluar")
		fmt.Println("==================")
		fmt.Print("Pilih: ")

		fmt.Scan(&pilihan)

		switch pilihan {
		case 1: 
			fmt.Printf("\nSaldo Anda: Rp.%d\n", saldo)
		case 2:
			fmt.Print("Masukkan jumlah uang yang ingin ditarik: Rp.")
			fmt.Scan(&jumlah)

			if jumlah <= saldo {
				saldo -= jumlah
				fmt.Println("\nAnda berhasil menarik uang sebesar Rp", jumlah)
				fmt.Println("Saldo tersisa: Rp", saldo)
			} else {
				fmt.Printf("\nSaldo tidak mencukupi\n")
			}
		case 3:
			fmt.Printf("\nKeluar dari sistem. Terima kasih!\n")
		default:
			fmt.Println("\nPilihan tidak valid. Masukkan angka 1 atau 2 atau 3!")
		}
	}
}
