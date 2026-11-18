package main 
import "fmt"

const NMAX int = 1000
type penduduk struct {
	nama string
	tahunLahir int
	kotaAsal string
}

type tabPenduduk [NMAX]penduduk

func cariTahun(T tabPenduduk, n, th int, idx *int) {
	*idx = -1
	for i := 0; i < n; i++ {
		if T[i].tahunLahir == th {
			*idx = i
			break
		}
	}
}

func tertuaTermuda(T tabPenduduk, n int, tua, muda *penduduk) {
	*tua = T[0]
	*muda = T[0]

	// tertua termuda
	for i := 0; i < n; i++ {
		if T[i].tahunLahir < tua.tahunLahir {
			*tua = T[i]
		}
	
		if T[i].tahunLahir > muda.tahunLahir {
			*muda = T[i]
		} 
	}
}

func main () {
	var (
		T tabPenduduk
		tua, muda penduduk
		n,tahunLahir,idx int
	)
	
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].nama, &T[i].tahunLahir, &T[i].kotaAsal)
	}	

	// Cari tahun
	fmt.Scan(&tahunLahir)
	cariTahun(T, n, tahunLahir, &idx)

	if idx != -1 {
		fmt.Printf("%s %d %s Ditemukan di index ke-%d\n", T[idx].nama, T[idx].tahunLahir, T[idx].kotaAsal, idx)
	} else {
		fmt.Println("Data tidak ditemukan")
	}

	// Sorting penduduk tertua & termuda
	tertuaTermuda(T, n, &tua, &muda)
	fmt.Printf("Penduduk Tertua: %s (%d)\n", tua.nama, tua.tahunLahir)
	fmt.Printf("Penduduk Termuda: %s (%d)\n", muda.nama, muda.tahunLahir)
}