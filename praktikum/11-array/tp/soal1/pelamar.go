package main

import "fmt"

const NMAX int = 100
type applicant struct {
	id, name     string
	testGrade    float64
	testDuration int
}

type applicants [NMAX]applicant

func bacaPelamar(A *applicants, n *int) {
	var id string
	*n = 0
	fmt.Scan(&id)
	for id != "END" && *n < NMAX {
		A[*n].id = id
		fmt.Scan(&A[*n].name, &A[*n].testGrade, &A[*n].testDuration)
		*n++
		fmt.Scan(&id)
	}
}

func cetakPelamar(A applicants, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %.1f %d\n", A[i].id, A[i].name, A[i].testGrade, A[i].testDuration)
	}
}

func salinLolos(A applicants, n int, batas float64, B *applicants, m *int) {
	*m = 0
	for i := 0; i < n; i++ {
		if A[i].testGrade >= batas {
			B[*m] = A[i]
			*m++
		}
	}
}

func rataDurasi(B applicants, m int) float64 {
	if m == 0 {
		return 0
	}
	total := 0
	for i := 0; i < m; i++ {
		total += B[i].testDuration
	}
	return float64(total) / float64(m)
}

func main() {
	var data, lolos applicants
	var n, m int
	var batas float64

	// Input data pelamar
	bacaPelamar(&data, &n)
	fmt.Scan(&batas)

	salinLolos(data, n, batas, &lolos, &m)

	fmt.Println()
	fmt.Println("Data awal:")
	cetakPelamar(data, n)

	if m == 0 {
		fmt.Println("Tidak ada pelamar yang lolos")
	} else {
		fmt.Println()
		fmt.Printf("Pelamar lolos (nilai >= %.1f):\n", batas)
		cetakPelamar(lolos, m)

		fmt.Println()
		fmt.Println("Jumlah pelamar lolos:", m)
		fmt.Printf("Rata-rata durasi: %.2f menit\n", rataDurasi(lolos, m))
	}
}