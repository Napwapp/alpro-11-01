package main
import "fmt"

func main () {
	var (tahun,bulan,minggu,hari int)

	// Input hari
	fmt.Println("Masukkan jumlah hari: ")
	fmt.Scan(&hari)

	tahun = hari / 360
	hari %= 360
	bulan = hari / 30
	hari %= 30
	minggu = hari / 7
	hari %= 7

	fmt.Printf("\n%d Tahun, %d Bulan, %d Minggu, %d Hari", tahun, bulan, minggu, hari)
}
