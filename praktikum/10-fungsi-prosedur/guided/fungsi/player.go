package main

import "fmt"

const NMAX int = 999

type player struct {
	nickname string
	score    int
}

type tabPlayer [NMAX]player

func rataRata(T tabPlayer, n int) float64 {
	total := 0
	for i := 0; i < n; i++ {
		total += T[i].score
	}
	return float64(total) / float64(n)
}

func idxTertinggi(T tabPlayer, n int) int {
	idx := 0
	for i := 1; i < n; i++ {
		if T[i].score > T[idx].score {
			idx = i
		}
	}
	return idx
}

func hitungDiAtas(T tabPlayer, n int, batas float64) int {
	jumlah := 0
	for i := 0; i < n; i++ {
		if float64(T[i].score) > batas {
			jumlah++
		}
	}
	return jumlah
}

func main() {
	var (
		T tabPlayer
		n int
	)

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].nickname, &T[i].score)
	}

	rata := rataRata(T, n)
	best := idxTertinggi(T, n)

	fmt.Println()
	fmt.Printf("Rata-rata skor: %.2f\n", rata)
	fmt.Printf("Player terbaik: %s (%d)\n", T[best].nickname, T[best].score)
	fmt.Printf("Jumlah player di atas rata-rata: %d\n", hitungDiAtas(T, n, rata))
}