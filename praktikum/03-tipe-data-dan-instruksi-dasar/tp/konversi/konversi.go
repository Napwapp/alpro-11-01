package main
import "fmt"

func main() {
	var(
		mil, km float64
	)

	fmt.Println("Konversi mil ke km. Masukkan mil:")
	fmt.Scan(&mil)
	km = mil * 1.6

	fmt.Printf("%.1f", km)
}
