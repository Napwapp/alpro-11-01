package main
import "fmt"

func main() {
	var jumlahA, jumlahB int
	jumlahA, jumlahB = 0, 0
	pemenang := ""
	dibaca := 0

	for dibaca < 10 {
		var s string
		_, err := fmt.Scan(&s)
		if err != nil {
			break
		}

		for _, c := range s {
			dibaca++
			if c == 'A' {
				jumlahA++
				if jumlahA == 4 && pemenang == "" {
					pemenang = "A"
				}
			} else if c == 'B' {
				jumlahB++
				if jumlahB == 4 && pemenang == "" {
					pemenang = "B"
				}
			}
		}
	}

	if jumlahA != 5 || jumlahB != 5 {
		fmt.Println("tidak valid")
	} else {
		fmt.Println(pemenang)
	}
}
