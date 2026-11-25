package main
import "fmt"

func add(a,b int32) int32 {
	return a + b
}

// 
func eating() {
	fmt.Println("Hello World!")
	fmt.Println("Hello World!")
	fmt.Println("Hello World!")
	fmt.Println("Hello World!")
}

// Prosedur
func hitungF(x, y int) {
	var hasil int
	
	hasil = 2*x + y - 5
	fmt.Println(hasil)
}

// Fungsi
func hitungFF(x, y int) int {
	var hasil int
	
	hasil = 2*x + y - 5
	return hasil
}

func main() {
	x := 2
	y := 3
	fmt.Println(add(2, 3))
	eating()

	hitungF(x, y) // prosedur
	fmt.Println(hitungFF(x, y)) // fungsi
}