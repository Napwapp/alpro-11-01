package main

import "fmt"

func main() {
	var (
		jenis string
		jam,tarif   int
	)

	fmt.Print("Masukkan jenis kendaraan (Mobil/Motor): ")
	fmt.Scan(&jenis)

	fmt.Print("Masukkan lama parkir (jam): ")
	fmt.Scan(&jam)

	if jam < 1 {
		fmt.Println("Lama parkir tidak valid")
		return
	}

	switch jenis {
	case "Mobil", "mobil":
		tarif = 5000 + (jam-1)*3000
		if jam > 24 {
			tarif += 50000
		}
	case "Motor", "motor":
		tarif = 3000 + (jam-1)*2000
		if jam > 24 {
			tarif += 30000
		}
	default:
		fmt.Println("Jenis kendaraan tidak valid")
		return
	}

	fmt.Println("Tarif parkir:", tarif)
}
