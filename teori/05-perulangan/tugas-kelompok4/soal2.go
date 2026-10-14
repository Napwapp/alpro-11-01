package main

import "fmt"

func main() {
	var transaksi, saldo int

	fmt.Scan(&transaksi)

	for transaksi != 0 && saldo+transaksi >= 0 {
		saldo += transaksi
		fmt.Scan(&transaksi)
	}

	fmt.Print(saldo)
}