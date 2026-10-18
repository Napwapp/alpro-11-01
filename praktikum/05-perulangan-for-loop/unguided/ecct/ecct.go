package main

import "fmt"

func main() {
	var (
		i, m, p, v, ga, fs, ci, tc, mahasiswaLulus int
		avg float64
		hasil []string
	)

	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&m)


	for i = 1; i <= m; i++ {
		fmt.Printf("Masukkan nilai mahasiswa %d (1-4)\n", i)

		// Rentang nilai 1-4
		fmt.Print("P: "); fmt.Scan(&p)
		fmt.Print("V: "); fmt.Scan(&v)
		fmt.Print("GA: "); fmt.Scan(&ga)
		fmt.Print("FS: "); fmt.Scan(&fs)
		fmt.Print("CI: "); fmt.Scan(&ci)
		fmt.Print("TC: "); fmt.Scan(&tc)

		// total dan avg
		totalNilai := p + v + ga + fs + ci + tc
		avg = float64(totalNilai) / 6

		if avg >= 3.33 {
			hasil = append(hasil, "Ya")
			mahasiswaLulus++
		} else {
			hasil = append(hasil, "Tidak")
		}
	}

	for _, sertifikat := range hasil {
		fmt.Println(sertifikat)
	}

	fmt.Print(mahasiswaLulus)
}
