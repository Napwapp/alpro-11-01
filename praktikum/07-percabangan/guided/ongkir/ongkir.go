package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		berat            float64
		zona, tarifPerKg int
		layanan          string
	)

	fmt.Scan(&berat, &zona, &layanan)
	if berat > 30 {
		fmt.Println("Berat melebihi batas maksimum (30kg)")
		return
	}

	switch zona {
	case 1:
		tarifPerKg = 10000
	case 2:
		tarifPerKg = 15000
	case 3:
		tarifPerKg = 25000
	default:
		fmt.Println("Zona tidak valid!")
		return
	}

	beratBulat := int(math.Ceil(berat))
	total := beratBulat * tarifPerKg

	switch layanan {
	case "Express":
		total += total / 2
	case "Reguler":
		total += 0
	default:
		fmt.Println("Layanan tidak valid")
		return
	}

	fmt.Printf("%d\n%d", beratBulat, total)
}
