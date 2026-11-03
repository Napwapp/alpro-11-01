package main
import "fmt"

func main() {
	total := 0

	for hari := 1; hari <= 5; hari++ {
		x := -1 // agar perulangan di bawah berjalan minimal sekali
		for x < 0 || x > 200 {
			fmt.Printf("Hari %d : ", hari)
			_, err := fmt.Scan(&x)
			if err != nil {
				return
			}
		}
		total += x
	}

	fmt.Printf("Jumlah pengunjung : %d orang\n", total)
}
