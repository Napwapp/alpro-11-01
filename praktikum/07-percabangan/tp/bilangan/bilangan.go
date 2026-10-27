package main
import "fmt"

func main() {
	var bil int
	
	fmt.Print("Masukkan Bilangan: ")
	fmt.Scan(&bil)

	if bil == 0 {
		fmt.Print("Nol")
	} else if bil % 2 == 00 {
		fmt.Print("Genap")

		if bil > 0 {
			fmt.Println(", Positif")
		} else {
			fmt.Println(", Negatif")
		}
	} else if bil % 2 != 0 {
		fmt.Print("Ganjil")

		if bil > 0 {
			fmt.Println(", Positif")
		} else {
			fmt.Println(", Negatif")
		}
	} else {
		fmt.Print("Bilangan yang dimasukkan tidak valid")
	}
}