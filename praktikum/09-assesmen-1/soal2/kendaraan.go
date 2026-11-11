package main

import "fmt"

func main() {
	var (
		tipe                   string
		totalm, totalk, totalt int
	)

	for {
		fmt.Scan(&tipe)

		if tipe != "M" && tipe != "K" && tipe != "T" {
			break
		}

		switch tipe {
		case "M":
			totalm++
		case "K":
			totalk++
		case "T":
			totalt++
		}
	}

	fmt.Printf("Tipe M: %d\nTipe K: %d\nTipe T: %d\n", totalm, totalk, totalt)
}
