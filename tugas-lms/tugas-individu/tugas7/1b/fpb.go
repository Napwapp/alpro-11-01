package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	batas := n
	if m < n {
		batas = m
	}

	for i := batas; i >= 1; i-- {
		if n%i == 0 && m%i == 0 {
			fmt.Println(i)
			break
		}
	}
}
