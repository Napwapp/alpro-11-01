package main

import "fmt"

func main() {
	var (
		penghasilan, pajak float64
	)

	fmt.Print("Masukkan penghasilan (juta): ")
	fmt.Scan(&penghasilan)

	switch {
	case penghasilan <= 50:
		pajak = penghasilan * 5.0 / 100.0
		fmt.Printf("Total pajak: %.1f juta\n", pajak)

	case penghasilan > 50 && penghasilan <= 100:
		pajak = (50 * 5.0 / 100.0) + ((penghasilan - 50) * 10.0 / 100.0)
		fmt.Printf("Total pajak: %.1f juta\n", pajak)

	case penghasilan > 100 && penghasilan <= 200:
		pajak = (50 * 5.0 / 100.0) + (50 * 10.0 / 100.0) + ((penghasilan-100)*15.0/100.0)
		fmt.Printf("Total pajak: %.1f juta\n", pajak)

	case penghasilan > 200:
		pajak = (50 * 5 / 100) + (50 * 10 / 100) + (100 * 15 / 100) + ((penghasilan-200)*20/100)
		fmt.Printf("Total pajak: %.1f juta\n", pajak)
	}
}
