package main

import "fmt"

func main() {
	var (
		a, b int
	)

	// Input
	fmt.Println("Masukkan 2 angka untuk menampilkan hasil operasi dari kedua angka tersebut (+, -, *, /, %) ")
	fmt.Println("Masukkan angka pertama: ")
	fmt.Scanln(&a)

	for b == 0 {
		fmt.Println("\nMasukkan angka kedua: ")
		fmt.Scanln(&b)

		fmt.Println("Angka kedua tidak boleh 0. Silahkan coba lagi! ")
	}
	
	// Output menggunakan formatting
	// % merupakan placeholder yang akan menampilkan nilai dari argumen setelahnya sesuai urutan
	// d merupakan verb integer yang
	fmt.Println("Hasil :")
	fmt.Printf("\n%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
	fmt.Printf("%d %% %d = %d\n", a, b, a%b)	
}
