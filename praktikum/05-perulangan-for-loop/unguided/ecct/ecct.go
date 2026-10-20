package main

import "fmt"

// Function untuk range nilai 1-4
func bacaNilai(label string) int {
	var nilai int

	for {
		fmt.Printf("%s: ", label)
		fmt.Scan(&nilai)

		// Rentang nilai 1-4
		if nilai >= 1 && nilai <= 4 {
			return nilai
		}

		fmt.Println("Masukkan nilai pada rentang 1-4. Coba lagi")
	}
}

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

		// panggil func bacaNilai untuk nilai supaya rentang 1-4 saja
		p = bacaNilai("P")
		v = bacaNilai("V")
		ga = bacaNilai("GA")
		fs = bacaNilai("FS")
		ci = bacaNilai("CI")
		tc = bacaNilai("TC")

		// total dan avg
		totalNilai := p + v + ga + fs + ci + tc
		avg = float64(totalNilai) / 6

		fmt.Printf("Nilai rata-rata mahasiswa ke-%d: %.2f\n\n", i, avg)
		if avg >= 3.33 {
			hasil = append(hasil, "Mendapatkan Sertifikat")
			mahasiswaLulus++
		} else {
			hasil = append(hasil, "Tidak Mendapatkan Sertifikat")
		}
	}

	for _, sertifikat := range hasil {
		fmt.Println(sertifikat)
	}

	fmt.Printf("%d Mahasiswa mendapatkan sertifikat", mahasiswaLulus)
}
