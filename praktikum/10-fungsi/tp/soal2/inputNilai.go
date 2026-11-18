package main
import "fmt"

const NMAX int = 100

type student struct {
	nim, name string
	grade float64
}

type students [NMAX]student

func cariNIM(T students, n int, nim string) int {
	for i := 0; i < n; i++ {
		if T[i].nim == nim {
			return i
		}
	}
	return -1
}

func rataRata(T students, n int) float64 {
	var total float64

	if n == 0 {
		return 0
	}
	
	for i := 0; i < n; i++ {
		total += T[i].grade
	}
	return total / float64(n)
}

func main() {
	var (
		T students
		nim, name string
		grade float64
		n int
	)

	for {
		_, err := fmt.Scan(&nim)
		if err != nil || nim == "STOP" {
			break
		}

		fmt.Scan(&name, &grade)

		if cariNIM(T, n, nim) == -1 && n < NMAX {
			T[n] = student{nim, name, grade}
			n++
		}
	}
	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Println(T[i].nim, T[i].name, T[i].grade)
	}

	fmt.Printf("Rata-rata nilai: %.2f", rataRata(T, n))
}