package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	terbesar := n
	if m > n {
		terbesar = m
	}

	kpk := terbesar
	for kpk%n != 0 || kpk%m != 0 {
		kpk += terbesar
	}
	fmt.Println(kpk)
}
