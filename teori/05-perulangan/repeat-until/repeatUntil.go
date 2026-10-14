package main

import "fmt"

func main() {
	var i,n int

	i = 1

	fmt.Scan(&n)
	for {
		fmt.Println(i)
		if i == n {
			break
		}
		i++
	}
}