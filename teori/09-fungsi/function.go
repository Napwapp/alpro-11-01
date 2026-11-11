package main
import "fmt"

func hitung(x, y int) int {
	return 2*x + y - 5
}

func main() {
	fmt.Println(hitung(20, 2))
	fmt.Println(hitung(15, 2))
	fmt.Println(hitung(144, 2))
}
