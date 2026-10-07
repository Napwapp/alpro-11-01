package main
import "fmt"

func main() {
	var (
		makanan []string
		harga []int
		pilihan int
	)

	harga = []int{
		15000,
		5000,
		20000,
		10000,
	}

	makanan = []string{
		"Nasi Goreng",
		"Mi Goreng",
		"Sate",
		"Gado-Gado",
	}

	fmt.Println("============= Makanan =============")
	fmt.Println("Pilih menu makanan")
	for i, m := range makanan {
		fmt.Printf("%d. %s\n", i+1, m)
	}

	fmt.Print("Masukkan pilihan (1-4): ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		fmt.Printf("%s\nHarga: Rp %d\n", makanan[0], harga[0])
	case 2:
		fmt.Printf("%s\nHarga: Rp %d\n", makanan[1], harga[1])
	case 3:
		fmt.Printf("%s\nHarga: Rp %d\n", makanan[2], harga[2])
	case 4:
		fmt.Printf("%s\nHarga: Rp %d\n", makanan[3], harga[3])
	default:
		fmt.Println("Pilihan tidak valid")
	}
}
