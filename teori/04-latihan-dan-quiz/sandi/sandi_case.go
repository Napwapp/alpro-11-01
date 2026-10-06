// pemecah kode
package main

import "fmt"

func main() {
	var jumlah, n, d1, d2, total int

	fmt.Scan(&jumlah)
	total = 0

	for i := 1; i <= jumlah; i++ {
		fmt.Scan(&n)
		d1 = n / 1000
		d2 = n % 10
		total += d1 + d2
	}

	fmt.Println(total)
}
