package main

import "fmt"

func main() {
	var (
		n, faktor int
		is_prime  bool
	)
	fmt.Println("Masukkan Bilangan:")
	fmt.Scan(&n)

	if n < 1 {
		fmt.Println("Masukkan minimal angka 1")
		return
	}

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			faktor = i
			fmt.Printf("%d ", faktor)
		}

		if n%2 != 0 && n%i == 0 {
			is_prime = true
		} else {
			is_prime = false
		}
	}

	if is_prime == true {
		fmt.Printf("\n%d Merupakan bilangan prima", n)
	} else {
		fmt.Printf("\n%d Bukan bilangan prima", n)
	}
}
