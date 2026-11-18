package main

import (
	"fmt"
	"unicode"
)

const NMAX = 1001

type murid struct {
	nama string
}
type tabMurid [NMAX]murid

func cariMurid(T tabMurid, n int, nama string, posisi *int, ketemu *bool) {
	/*
		I.S. T[0..n-1] dan nama terdefinisi
		F.S. ketemu bernilai true dan posisi berisi urutan absen (mulai dari 1)
		jika nama ditemukan; jika tidak, ketemu bernilai false
	*/

	*ketemu = false
	for i := 0; i < n; i++ {
		if T[i].nama == nama {
			*ketemu = true
			*posisi = i + 1
			return
		}
	}

}

func hitungHuruf(T tabMurid, n int, h byte, jumlah *int) {
	/* 
		IS : T[0..n-1] dan h terdefinisi
		FS : jumlah berisi banyaknya nama pada T yang diawali huruf h
	*/

	for i := 0; i < n; i++ {
		if T[i].nama != "" && T[i].nama[0] == h {
			*jumlah++ 
			return
		}
	}
}

func main() {
	var (
		T                 tabMurid
		n, posisi, jumlah int
		nama              string
		ketemu            bool
		h                 byte
	)

	// Daftar nama sejumlah n siswa
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].nama)
	}

	// Cari murid
	fmt.Scan(&nama)
	cariMurid(T, n, nama, &posisi, &ketemu)
	if ketemu {
		fmt.Printf("Murid terdaftar dan berada di urutan absen ke-%d\n", posisi)
	} else {
		fmt.Println("Murid tidak terdaftar")
	}

	// Hitung huruf
	h = nama[0]
	hitungHuruf(T, n, h, &jumlah)
	fmt.Printf("Jumlah murid dengan huruf awal %c: %d", unicode.ToUpper(rune(h)), jumlah+1)
}
