package main
import "fmt"

func hitungLuas(alas, tinggi int, luas *float64) {
	var n int

	fmt.Scan(&n)
	daftarLuas := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&alas, &tinggi)

		*luas = float64(alas) * float64(tinggi) / 2
		daftarLuas[i] = *luas
	}

	fmt.Println()

	for _,hasil := range daftarLuas {
		fmt.Println(hasil)
	}
}

func main() {
	var (
		alas, tinggi int
		luas float64
	)
	
	hitungLuas(alas, tinggi, &luas)
}
