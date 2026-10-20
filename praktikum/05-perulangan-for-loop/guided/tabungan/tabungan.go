package main
import "fmt"

func main() {
	var tabunganHariIni, b, n, total int

	fmt.Scan(&tabunganHariIni)
	fmt.Scan(&b)
	fmt.Scan(&n)

	for i := 1; i <= n; i++  {
		fmt.Printf("Hari ke %d: %d\n", i, tabunganHariIni)
		total = total + tabunganHariIni
		tabunganHariIni = tabunganHariIni + b
	}

	fmt.Print("Total: ", total)
}