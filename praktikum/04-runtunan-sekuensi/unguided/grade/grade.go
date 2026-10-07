package main
import "bufio"
import "fmt"
import "os"

func main() {
	var (
		nilai float64
		nama,grade string
	)
	// bufio scanner untuk input
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("Masukkan nama: ")
	scanner.Scan()
	nama = scanner.Text()

	fmt.Println("Masukkan nilai: ")
	fmt.Scan(&nilai)

	switch {
	case nilai >= 90 && nilai <= 100:
		grade = "A"
	case nilai >= 80 && nilai < 90:
		grade = "B"
	case nilai >= 70 && nilai < 80:
		grade = "C"
	case nilai >= 60 && nilai < 70:
		grade = "D"
	default:
		grade = "E"
	}

	fmt.Printf("Nama: %s\n Grade: %s", nama, grade)
}