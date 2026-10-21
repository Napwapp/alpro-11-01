# <h1 align="center">Tugas Pendahuluan Modul 06 - Perulangan While</h1>

<p align="center">Muhamad Nawaf Abduh - 109092630003</p>

### 1. lulus.go

```go
package main

import "fmt"

func main() {
	var nilai int

	for {
		fmt.Println("Masukkan Nilai: ")
		fmt.Scan(&nilai)

		if nilai == -1 {
			fmt.Println("Program Selesai")
			break
		}

		if nilai >= 60 && nilai <= 100 {
			fmt.Println("Anda lulus ujian")
		} else if nilai < 60 && nilai >= 0{
			fmt.Println("Anda tidak lulus ujian")
		} else {
			fmt.Println("Nilai Tidak Valid")
		}
	}
}
```

##### Output

![Screenshot output lulus.go](/praktikum/06-perulangan-while-loop-dan-repeat/tp/lulus/output.png)

#### Deskripsi
`lulus.go` Merupakan program untuk memasukkan `nilai` secara berulang. Jika `nilai` yang dimasukkan >= 60, akan menampilkan pesan "Anda lulus ujian" jika < 60 sampai 0 akan menampilkan pesan "Nilai Tidak Valid" Progam ini akan terus meminta pengguna unutk memasukkan nilai hingga pengguna memasukkan angka `-1` untuk menghentikan program dan akan menampilkan pesan "Program Selesai".

### 2. deret.go

```go
package main
import "fmt"

func main() {
	var n, total int 

	i := 1
	fmt.Print("Masukkan Bilangan: ")
	fmt.Scan(&n)

	for i <= n {
		fmt.Print(i, " ")
		total = total + i
		i++
	}

	fmt.Printf("\nTotal: %d",total)
}
```

##### Output

![Screenshot output total bilangan deret](/praktikum/06-perulangan-while-loop-dan-repeat/tp/deret/output.png)

#### Deskripsi
`deret.go` Merupakan program untuk mencari deret dari bilangan `n` yang dimasukkan, lalu menjumlahkan semua bilangan deret nya. Contoh:
`n = 5` maka `1 + 2 + 3 + 4 + 5 = 15`

### 3. 2terendah.go


```go
package main

import "fmt"

func main() {
	var n, nilai int
	var nama string

	var lowest, secondLowest string
	min1 := 999999
	min2 := 999999

	fmt.Scan(&n)

	i := 1
	for i <= n {
		fmt.Print("Nama:")
		fmt.Scan(&nama)

		fmt.Print("Nilai:")
		fmt.Scan(&nilai)

		if nilai < min1 {
			min2 = min1
			fmt.Print(min2)
			secondLowest = lowest
			fmt.Print(secondLowest)

			min1 = nilai
			fmt.Print(min1)
			lowest = nama
		} else if nilai < min2 {
			min2 = nilai
			fmt.Print(min2)
			secondLowest = nama
		}

		i++
	}

	fmt.Println("Siswa dengan Nilai Terendah:", lowest)
	fmt.Println("Siswa dengan Nilai Terendah ke-2:", secondLowest)
}
```

##### Output

![Screenshot output 2 nilai terendah](/praktikum/06-perulangan-while-loop-dan-repeat/tp/dua-terendah/output.png)

#### Deskripsi
`2terendah.go` Merupakan program untuk mencari 2 nilai terendah dari sejumlah `n` siswa yang dimasukkan, lalu menampilkan nama siswa yang memiliki nilai terendah pertama dan terendah ke-2.

1. Inisialisasi `lowest`, `secondLowest` untuk menampung nama siswa. Nilai `min1`, `min2` diinisialisasi dengan angka yang sangat besar yaitu `999999` sebagai nilai awal pembanding untuk mencari nilai minimum supaya `nilai` < `min1` atau `min2`.
2. Di dalam perulangan, program meminta input `nama` dan `nilai`, lalu dilakukan pengecekan kondisi:
   - `if nilai < min1`: Jika `nilai` yang dimasukkan lebih kecil dari `min1`, maka data terendah pertama sebelumnya digeser ke posisi 2 `min2 = min1` dan `secondLowest = lowest`. Lalu, `min1` dan `lowest` diperbarui dengan `nilai` dan `nama` yang baru saja dimasukkan.
   - `else if nilai < min2`: Jika nilainya tidak lebih kecil dari `min1` tetapi lebih kecil dari `min2`, maka hanya `nilai` dan `nama` siswa terendah ke-2 yang diperbarui `min2 = nilai` dan `secondLowest = nama`,

Setelah perulangan selesai, program akan menampilkan output nama siswa dengan nilai terendah (`lowest`) dan nilai terendah ke-2 (`secondLowest`).

### 4. biaya.go
```go
package main

import "fmt"

func main() {
	var (
		n int
		biayaTahunIni, p6 float64
	)

	biayaTahunIni = 10000000
	i := 2
	fmt.Scan(&n)

	fmt.Printf("Tahun ke 1: %0.f\n", biayaTahunIni)
	for i <= n {
		p6 = biayaTahunIni * 0.06
		biayaTahunIni = biayaTahunIni + p6
		fmt.Printf("Tahun ke %d: %0.f\n", i, biayaTahunIni)
		i++
	}

	fmt.Printf("Biaya tahun ke %d: %0.f",n, biayaTahunIni)
}
```

##### Output

![Screenshot output hitung biaya kuliah pada tahun ke n](/praktikum/06-perulangan-while-loop-dan-repeat/tp/biaya/output.png)

#### Deskripsi
`biaya.go` Merupakan program untuk membantu sebuah universitas menghitung perkiraan biaya kuliah dari tahun ke tahun. Biaya kuliah awalnya sebesar `10.000.000` dan setiap tahunnya akan bertambah `6%` dari biaya tahun sebelumnya hingga mencapai tahun ke-`n` yang dimasukkan.

1. `biayaTahunIni = 10000000` di inisialisasikan sebagai biaya kuliah pada tahun ke-1. 

`i := 2` i bernilai 2 supaya perulangan mulai menghitung biaya tahun ke 2 hingga tahun ke-`n`. 

   - `p6 = biayaTahunIni * 0.06`: Menghitung kenaikan biaya sebesar 6% dari biaya tahun sebelumnya.
   - `biayaTahunIni = biayaTahunIni + p6`: Menambahkan kenaikan 6% tersebut ke biaya kuliah saat ini.
   - Menampilkan biaya untuk tahun ke-`i`.
   - `i++`: Menaikkan nilai variabel `i` sebesar 1 di setiap akhir perulangan.

Setelah perulangan selesai, program akan menampilkan ringkasan biaya kuliah pada tahun ke-`n`.

## Kesimpulan
Pendekatan `while` di perulangan pada bahasa pemrograman Go sebenarnya konsepnya hampir sama dengan perulangan yang menggunakan 3 komponen (`inisialisasi`, `kondisi`, `post statement`). Bedanya, ada pada penempatan komponennya, di mana inisialisasi variabel dilakukan di luar blok perulangan, pada kata kunci `for` nya hanya memiliki kondisi, dan penambahan atau pembaruan nilainya (post statement) dijalankan di dalam blok perulangan. 

Kode didalam perulangan tersebut hanya akan terus dieksekusi selama kondisinya masih bernilai `true`.

Selain itu, ada juga bentuk perulangan dengan pendekatan `repeat-until`.

Kali ini, `for` tidak memiliki kondisi 
```go
for {
	...

	if kondisi {
		break
	}
}
```
Perulangan ini pasti akan dijalankan terlebih dahulu lalu program akan mengecek kondisi yang memiliki `break` didalam perulangan jika kondisi tersebut bernilai `true` maka perulangan akan berhenti, jika `false` perulangan akan terus dijalankan sampai kondisi tersebut bernilai `true`

Sehingga titik berhenti perulangannya ada di dalam blok perulangan itu sendiri  melalui pernyataan `if` dan `break` (jika tidak ada, maka akan terjadi yang namanya `infinite-loop` atau loop tanpa berhenti).
