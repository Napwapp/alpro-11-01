package main
import "fmt"

const NMAX int = 999

type gempa struct {
	nama      string
	tahun     int
	magnitudo float64
}

type tabGempa [NMAX]gempa

func kategori(m float64) string {
	var kategori string

	if m < 4.0 {
		kategori = "Minor"
		return kategori
	} else if m >= 4.0 && m < 6.0 {
		kategori = "Sedang"
		return kategori
	} else if m >= 6.0 && m < 7.0 {
		kategori = "Kuat"
		return kategori
	} else if m >= 7.0 {
		kategori = "Besar"
		return kategori
	} else {
		return "Tidak valid"
	}
}

// Hitung rata-rata magnitudo dari banyaknya gempa
func rataMagnitudo(T tabGempa, n int) float64 {
	var total float64
	for i := 0; i < n; i++ {
		total += T[i].magnitudo
	}
	return total / float64(n)
}

// Gempa terkuat
func idxTerkuat(T tabGempa, n int) int {
	idx := 0
	for i := 1; i < n; i++ {
		if T[i].magnitudo > T[idx].magnitudo {
			idx = i
		}
	}
	return idx
}

// Hitung banyaknya gempa besar
func hitungKategori(T tabGempa, n int, k string) int {
	jumlah := 0
	for i := 0; i < n; i++ {
		if kategori(T[i].magnitudo) == k {
			jumlah++
		}
	}
	return jumlah
}

func main() {
	var (
		T tabGempa
		n int
	)
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].nama, &T[i].tahun, &T[i].magnitudo)
	}
	fmt.Println()
	// daftar gempa beserta kategorinya
	for i := 0; i < n; i++ {
		fmt.Printf("%s %d %.1f %s\n", T[i].nama, T[i].tahun, T[i].magnitudo, kategori(T[i].magnitudo))
	}

	fmt.Println()
	fmt.Printf("Rata-rata magnitudo: %.2f\n", rataMagnitudo(T, n))

	k := idxTerkuat(T, n)
	fmt.Printf("Gempa terkuat: %s %d %.1f\n", T[k].nama, T[k].tahun, T[k].magnitudo)

	fmt.Printf("Jumlah gempa kategori Besar: %d\n", hitungKategori(T, n, "Besar"))
}