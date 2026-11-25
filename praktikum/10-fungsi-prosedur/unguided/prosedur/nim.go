package main

import "fmt"

const NMAX int = 2002
type tabNIM [NMAX]int

func cariNIM(T tabNIM, n, x int, idx *int) {
	*idx = -1
	i := 0
	// berhenti begitu ketemu, jadi kalau ada NIM kembar yang diambil yang pertama
	for i < n && *idx == -1 {
		if T[i] == x {
			*idx = i
		}
		i++
	}
}

func ubahNIM(T *tabNIM, idx, baru int) {
	T[idx] = baru
}

func cetak(T tabNIM, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(T[i])
	}
}

func main() {
	var (
		T tabNIM
		n int
		x int
		idx int
		baru int
	)

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i])
	}
	fmt.Scan(&x)

	cariNIM(T, n, x, &idx)

	if idx == -1 {
		fmt.Println()
		fmt.Println("NIM tidak ditemukan dalam sistem")
	} else {
		// NIM pengganti hanya dibaca kalau NIM x ketemu
		fmt.Scan(&baru)
		ubahNIM(&T, idx, baru)

		fmt.Println()
		fmt.Printf("Mahasiswa dengan NIM %d terdaftar pada posisi ke-%d\n", x, idx+1)
		fmt.Println("Daftar NIM setelah terjadi perubahan data:")
		cetak(T, n)
	}
}