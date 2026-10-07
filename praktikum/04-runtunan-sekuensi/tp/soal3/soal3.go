// Soal 3 – Menentukan Jumlah Hari dalam Sebulan Berdasarkan Tahun dan Bulan
package main

import "fmt"

func main() {
	var (
		tahun, hari  int
		bulan string
	)

	fmt.Print("Masukkan tahun: ")
	fmt.Scanln(&tahun)
	fmt.Print("Masukkan bulan dalam format (Jan, Feb, Mar, Apr, Mei, Jun, Jul, Agu, Sep, Okt, Nov, Des): ")
	fmt.Scanln(&bulan)

	switch bulan {
		case "Jan":
			hari = 31
		case "Feb":
			if (tahun%4 == 0 && tahun%100 != 0) || (tahun%400 == 0) {
				hari = 29
			} else {
				hari = 28
			}
		case "Mar":
			hari = 31
		case "Apr":
			hari = 30
		case "Mei":
			hari = 31
		case "Jun":
			hari = 30
		case "Jul":
			hari = 31
		case "Agu":
			hari = 31
		case "Sep":
			hari = 30
		case "Okt":
			hari = 31
		case "Nov":
			hari = 30
		case "Des":
			hari = 31
	default:
		fmt.Println("-")
		return
	}
	
	fmt.Println(hari)
}
