package main

import "fmt"

func main() {
	var (
		n, jumlahFaktor int
	)
	fmt.Println("Masukkan Bilangan:")
	fmt.Scan(&n)

	if n < 1 {
		fmt.Println("Masukkan minimal angka 1")
		return
	}

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Println(i)
			jumlahFaktor++
		}

		if jumlahFaktor == 2 {
			fmt.Printf("\n%d Merupakan bilangan prima", n)
		} else {
			fmt.Printf("\n%d Bukan bilangan prima", n)
		}
	}
}
