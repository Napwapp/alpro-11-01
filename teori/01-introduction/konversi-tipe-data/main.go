package main
import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 = int64(nilai32)
	var nilai16 = int16(nilai32)

	fmt.Println("Hasil konversi : ", nilai64)
	fmt.Println("Hasil konversi : ", nilai16)

	name := "Nawaf"
	firstLetter := name[0]
	firstLetterString := string(firstLetter)
	fmt.Println("Huruf pertama : ", firstLetterString)

	// Type declaration
	// Yaitu mengaliaskan tipe data baru dari tipe data yang sudah ada.
	type (
		noKtp string
		alamat string	
	)

	var ktp noKtp = "1234567890"
	var address alamat = "Jl. Mawar No. 123"
	fmt.Println("No KTP : ", ktp)
	fmt.Println("Alamat : ", address)
}
