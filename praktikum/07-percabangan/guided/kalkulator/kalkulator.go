package main

import "fmt"

func main() {
	var (
		op string
		a, b int
	)

	fmt.Scan(&a, &op, &b)

	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Tidak bisa dibagi dengan nol")
		} else {
			fmt.Printf("%.2f ", float64(a) / float64(b))
		} 
	case "%":
		if b == 0 {
			fmt.Println("Tidak bisa dibagi dengan nol")
		} else {
			fmt.Println(a%b)
		}
	default:
		fmt.Println("Operator tidak valid")
		return
	}
}
