package main

import "fmt"

func main() {
	var (
		totalBelanja, totalAkhir float64
		mahasiswaAssesmen bool
	)

	fmt.Scan(&totalBelanja, &mahasiswaAssesmen)
	if mahasiswaAssesmen == true {
		totalAkhir = totalBelanja * 0.35
		fmt.Println(totalAkhir)
	} else {
		fmt.Println("Tidak sedang assesmen")
	}
}
