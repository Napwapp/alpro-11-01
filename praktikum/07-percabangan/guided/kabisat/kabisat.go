package main
import "fmt"

func main() {
	var tahun int
	fmt.Scan(&tahun)

	if (tahun % 4 == 0 && tahun % 100 != 0) || (tahun % 400 == 0) {
		fmt.Printf("Ya, %d Merupakan tahun kabisat", tahun)
	} else {
		fmt.Printf("Tidak, %d Bukan tahun kabisat", tahun)
	}
}