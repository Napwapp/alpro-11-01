package main

import "fmt"

func main() {
	var nilai int

	for {
		fmt.Println("Masukkan Nilai: ")
		fmt.Scan(&nilai)

		if nilai == -1 {
			fmt.Println("Program Selesai")
			break
		}

		if nilai >= 60 && nilai <= 100 {
			fmt.Println("Anda lulus ujian")
		} else if nilai < 60 && nilai >= 0{
			fmt.Println("Anda tidak lulus ujian")
		} else {
			fmt.Println("Nilai Tidak Valid")
		}
	}
}
