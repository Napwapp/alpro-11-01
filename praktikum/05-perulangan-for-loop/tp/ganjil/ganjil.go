package main

import "fmt"

func main() {
	var n, total int
	fmt.Println("Masukkan bilangan:")
	fmt.Scan(&n)

	for i := 1; i <= n; i += 2 {
		fmt.Printf("%d ", i)
		total = total + i
	}
	fmt.Printf("\nTotal: ")
	fmt.Printf("%d\n", total)
}
