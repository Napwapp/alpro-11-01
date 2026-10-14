package main
import "fmt"

func main () {
	var a,b int

	fmt.Printf("Masukkan bilangan a, dan b:\na akan dikalikan mulai dari 1 sampai ke bilangan b\n")
	fmt.Scan(&a, &b)

	for i := 1; i <= b; i++ {
		hasilPerkalian := a * i
		fmt.Printf("%d x %d = %d\n", a, i, hasilPerkalian)
	}
}