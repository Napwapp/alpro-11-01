package main
import "fmt"

func main () {
	var celcius,raemur float64

	fmt.Println("Konversi suhu Celcius ke Raemur. Masukkan suhu Celcius: ")
	fmt.Scanln(&celcius)

	raemur = 4.0/5.0 * celcius

	fmt.Printf("\n%.2f", raemur)
}