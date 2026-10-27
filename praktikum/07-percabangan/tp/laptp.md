# <h1 align="center">Tugas Pendahuluan Modul 06 - Percabangan</h1>
<p align="center">Muhamad Nawaf Abduh - 109092630003</p>

### 1. Jelaskan perbedaan antara if, if-else, else if, dan switch pada bahasa pemrograman Go! Sebutkan juga dua perbedaan switch pada Go dibandingkan dengan switch pada Java.

#### Perbedaan if, if-else, else if, dan switch pada bahasa pemrograman Go
- `if` Blok kode didalam if akan dijalankan hanya jika suatu kondisi bernilai `true`, jika false blok kode di dalamnya tidak akan dijalankan 

- `if-else` Digunakan untuk menangani dua kemungkinan tindakan. Digunakan ketika terdapat kondisi yang bernilai `true` atau `false`. `if` untuk menjalankan tindakan/action saat kondisi nya bernilai `true` dan `else` untuk menjalankan tindakan/action saat kondisi nya bernilai `false`

- `else-if` Digunakan untuk memeriksa beberapa kondisi, program akan mengecek kondisinya satu per satu secara berurutan. Cocok digunakan ketika terdapat lebih dari dua kemungkinan kondisi tertentu pada program yang ingin dibuat. Begitu menemukan kondisi yang bernilai `true`, program akan menjalankan blok kode di dalam kondisi yang bernilai `true` dan melewati kondisi-kondisi berikutnya.

- `switch` (Pilih kasus). Fungsinya mirip dengan `else if` yang menjalankan blok kode pada suatu kondisi yang bernilai `true` (jika di switch itu `case`) namun dibandingkan `else-if`, penulisan `switch` lebih ringkas.

#### Perbedaan switch pada bahasa Go dengan switch pada Java

- `switch` di Java memerlukan `break` untuk mencegah program lanjut mengecek kondisi pada `case` lain walaupun `case` sebelumnya bernilai `true`. Sedangkan di Go, tidak memerlukan `break` karena sudah otomatis berhenti saat case yang bernilai `true` dieksekusi

- `switch` di Go dapat digunakan tanpa ekspresi, sehingga setiap `case` nya bisa berisi suatu kondisi nya langsung. Di Java, `switch` umumnya membandingkan nilai dengan pilihan `case`.


### 2. bilangan.go
```go
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
```

##### Output
![Screenshot output](/praktikum/07-percabangan/tp/bilangan/output.png)

#### Deskripsi
`bilangan.go` Merupakan program untuk mencari tahu bilangan yang dimasukkan adalah bilangan genap/ganjil dan bilangan positif/negatif, dan output nya akan "Nol" apabila bilangan yang dimasukkan adalah 0. Program ini mneggunakan percabangan dan kondisi untuk mengetahui bilangan yang dimasukkannya adalah bilangan genap/ganjil dan bilangan positif/negatif 

### 3. hari.go

```go
package main
import "fmt"

func main() {
	var hari int

	fmt.Print("Masukkan hari (1-7): ")
	fmt.Scan(&hari)

	switch hari {
	case 1:
		fmt.Println("Senin")
	case 2:
		fmt.Println("Selasa")
	case 3:
		fmt.Println("Rabu")
	case 4:
		fmt.Println("Kamis")
	case 5:
		fmt.Println("Jumat")
	case 6:
		fmt.Println("Sabtu")
	case 7:
		fmt.Println("Minggu")
	default:
		fmt.Println("Nomor hari tidak valid. Masukkan nomor (1-7)!")
	}
}
```

##### Output
![Screenshot output](/praktikum/07-percabangan/tp/hari/output.png)

#### Deskripsi
`hari.go` Merupakan program untuk menampilkan hari berdasarkan nomor/angka yang dimasukkan. 

1 = Senin
2 = Selasa
3 = Rabu
4 = Kammis
5 = Jumat
6 = Sabtu
7 = Minggu
Lainnya = Nomor hari tidak valid. Masukkan nomor (1-7)!

Program ini menggunakan percabangan dari `switch` yang mempunyai `case` nya masing-masing

### 4. terbesar.go
```go
package main
import "fmt"

func main () {
	var a, b, c int

	fmt.Println("Masukkan Bilangan a, b, dan c: ")
	fmt.Scan(&a, &b, &c)

	if a >= b && a >= c {
        fmt.Printf("\n%d", a)
    } else if b >= a && b >= c {
        fmt.Printf("\n%d", b)
    } else {
        fmt.Printf("\n%d", c)
    }
}
```

##### Output
![Screenshot output](/praktikum/07-percabangan/tp/terbesar/output.png)

#### Deskripsi
`terbesar.go` Merupakan program untuk mencari bilangan terbesar dari ketiga bilangan yang dimasukkan dengan memanfaatkan percabangan dan pengkondisian.

## Kesimpulan
Percabangan `if`, `else-if`, dan `switch` dapat membantu kita untuk membuat program yang tindakan/aksi nya hanya akan dijalankan sesuai dengan kondisi nya. Sehingga kita bisa membuat suatu tindakan/aksi yang berbeda pada suatu kondisi dengan kondisi yang lain, untuk untuk mengarahkan jalannya program agar mampu mengambil keputusan-keputusan berdasarkan data dan input yang diterima.