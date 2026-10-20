package main

import "fmt"

func main() {
	var (
		nilai float64
		grade string
	)

	fmt.Println("Masukkan nilai: ")
	fmt.Scan(&nilai)

	if nilai >= 85 && nilai <= 100 {
		grade = "A"
	} else if nilai >= 70 && nilai < 85 {
		grade = "B"
	} else if nilai >= 50 && nilai < 70 {
		grade = "C"
	} else if nilai < 50 && nilai == 0 {
		grade = "E"
	} else {
		fmt.Println("Nilai tidak valid")
	}

	fmt.Println(grade)
}
