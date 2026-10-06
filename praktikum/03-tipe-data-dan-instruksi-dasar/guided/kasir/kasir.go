package main

import "fmt"

func main () {
	var (x, lembar10rb, lembar5rb, lembar1rb int)

	// Input
	fmt.Println("Masukkan nominal uang: ")
	fmt.Scanln(&x)

	// Operasi
	lembar10rb = x / 10000 
	x %= 10000 
	lembar5rb = x / 5000
	x %= 5000
	lembar1rb = x / 1000

	// Output
	fmt.Println(lembar10rb)
	fmt.Println(lembar5rb)
	fmt.Println(lembar1rb)
}
