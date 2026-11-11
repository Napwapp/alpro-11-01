package main

import "fmt"

func digitTidakMenurun(n int) bool {
	sebelumnya := 10
	for n > 0 {
		digit := n % 10
		if digit > sebelumnya {
			return false
		}
		sebelumnya = digit
		n /= 10
	}
	return true
}

func jumlahDigit(n int) int {
	jumlah := 0
	for n > 0 {
		jumlah += n % 10
		n /= 10
	}
	return jumlah
}

func prima(n int) bool {
	if n < 2 {
		return false
	}
	for pembagi := 2; pembagi*pembagi <= n; pembagi++ {
		if n%pembagi == 0 {
			return false
		}
	}
	return true
}

func main() {
	var (
		// 0 <= l <= r
		l, r, jumlahBilPrima int
	)

	fmt.Scan(&l, &r)
	for i := l; ; i++ {
		if digitTidakMenurun(i) && prima(jumlahDigit(i)) {
			jumlahBilPrima++
		}
		if i == r {
			break
		}
	}
	fmt.Println(jumlahBilPrima)
}
