package main
import "fmt"

func main() {
	var a, b, n int

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&n)

	total := 0
	tabunganHariIni := a

	for i := 1; i <= n; i++  {
		total = total + tabunganHariIni
		tabunganHariIni = tabunganHariIni + b
		fmt.Println(tabunganHariIni)
	}

	fmt.Print(total)
}