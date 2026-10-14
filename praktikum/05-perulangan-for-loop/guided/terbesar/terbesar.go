package main

import "fmt"

func main() {
	var n, bilangan, bilanganTerbesar int

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&bilangan)

		if i == 1 || bilangan > bilanganTerbesar {
			bilanganTerbesar = bilangan
		}
	}

	fmt.Print(bilanganTerbesar)
}
