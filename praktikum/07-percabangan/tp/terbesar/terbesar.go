package main
import "fmt"

func main () {
	var a, b, c int

	fmt.Println("Masukkan Bilangan a, b, dan c: ")
	fmt.Scan(&a, &b, &c)

	if a >= b && a >= c {
        fmt.Printf("\n%d", a)
    } else if b >= a && b >= c {
        fmt.Printf("\n%d", b)
    } else {
        fmt.Printf("\n%d", c)
    }
}