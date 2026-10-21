package main
import "fmt"

func main() {
	var n, total int 

	i := 1
	fmt.Print("Masukkan Bilangan: ")
	fmt.Scan(&n)

	for i <= n {
		fmt.Print(i, " ")
		total = total + i
		i++
	}

	fmt.Printf("\nTotal: %d",total)
}
