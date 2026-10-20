# <h1 align="center">Laporan Praktikum Modul 05 - Perulangan For</h1>
<p align="center">Muhamad Nawaf Abduh - 109092630003</p>

## Dasar Teori

###  Perulangan (Looping)
Perulangan (Looping) Merupakan struktur kontrol yang mengeksekusi suatu blok kode secara berulang-ulang selama kondisi pada perulangan nya masih terpenuhi.

### Perulangan for-loop
Perulangan dalam bahasa Go dapat ditulis menggunakan kata kunci `for`. Tidak seperti pada bahasa pemrograman lain, Go hanya hanya bisa menggunakan `for` untuk perulangan. Sementara pada bebrapa bahasa pemrograman lain kita dapat menggunakan kata kunci lain seperti `while`, dan `do-while`
Namun, walaupun `while` dan `do-while` tidak disediakan pada bahas Go, kita tetap dapat menggunakan pendekatan yang sama hanya saja tetap menggunakan kata kunci `for` dengan penulisan kondisi pada perulangannya yang sedikit berbeda.

### Bentuk-bentuk perulangan yang bisa dituliskan dalam Go

#### 1. `for` dengan 3 komponen
- inisialiasi variabel
- kondisi
- Post Statement (untuk memperbarui nilai variabel setiap satu kali perulangan dijalankan)

```go
for inisialisasi; kondisi; postStatement {
    ...
}
```
Contoh Program:
```go
for i := 1; i <= 5; i++ {
    fmt.Printf("%d. Hello World", i)
}
```
Perulangan diatas akan menampilkan `Hello World` sebanyak 5 kali dengan nilai dari variabel `i`
```text
1. Hello World
2. Hello World
3. Hello World
4. Hello World
5. Hello World
```

#### 2. Perulangan hanya dengan kondisi
```go
for kondisi {
    ...
} 
```
Perulangan yang bentuknya hanya dengan kondisi, kode didalam perulangannya hanya akan dijalankan selama kondisi tersebut masih bernilai `true` (terpenuhi)


#### 3. Perulangan tanpa kondisi
```go
for {
    ...
}
```
Karena perulangan ini tidak memiliki pengkondisian apapun, dia akan dijalankan tersu menerus tanpa berhenti (loop). Perlu `if-break` untuk menghentikan perulangannya ketika telah mencapai kondisi tertentu sesuai dengan yang kita tentukan.

Contoh:
```go
i := 1
for {
    fmt.Println("Hello World")
    if i == 10 {
        break // Titik berhenti
    }

    i++
}
```

#### 4. For-range
For-range digunakan untuk melakukan iterasi terhadap suatu kumpulan data seperti `slice`, `array`, `string`, dan `map`. Setiap iterasi nya akan menghasilkan `index` dan `nilai` nya.

```go
buah := []string{"Apel", "Jeruk", "Mangga"}

for i, v := range buah {
    fmt.Println(buah)
}
```
Output:
```text 
0 Apel
1 Jeruk
2 Mangga
```

#### 5. Nested Loop (Perulangan Bersarang)
Perulangan juga dapat dituliskan lagi di dalam sebuah perulangan. 

```go
package main
import "fmt"

func main() {
    for i := 1; i <= 3; i++ {
        for j := 1; j <= 2; j++ {
            fmt.Print("* ")
        }
        fmt.Println("* ")
    }
}
```

Perulangan pertama `for i := 1; i <= 3; i++` akan menghasilkan Output:
```text
*
*
*
```

lalu perulangan yang ada didalam nya akan menghasilkan `for j := 1; j <= 2; j++` akan menghasilkan Output 2 `*` pada masing masing baris `*` sebelumnya.

```text
 * *
 * *
 * *
```

Sehingga hasil akhirnya adalah 
```text
* * *
* * *
* * *
```

## Guided
### 1.  tabungan.go
```go
package main
import "fmt"

func main() {
	var tabunganHariIni, b, n, total int

	fmt.Scan(&tabunganHariIni)
	fmt.Scan(&b)
	fmt.Scan(&n)

	for i := 1; i <= n; i++  {
		fmt.Printf("Hari ke %d: %d\n", i, tabunganHariIni)
		total = total + tabunganHariIni
		tabunganHariIni = tabunganHariIni + b
	}

	fmt.Print("Total: ", total)
}
```

#### Deskripsi
`tabungan.go` Merupakan program untuk menghitung total tabungan setelah `n` hari Dengan studi kasus: setiap hari, tabungannya bertambah sejumlah `b` Rupiah, lalu menghitung total nya setelah `n` Hari.

Input
```go
fmt.Scan(&tabunganHariIni) // Jumlah tabungan pada hari pertama dan yang akan ditambahkan dengan b
fmt.Scan(&b) // Penambahan tabungan setiap hari
fmt.Scan(&n) // Total hari
```

#### Contoh Input, Output
#### Input
```text
1000
500
5
```
#### Output
```text
Total: 10000
```
1000 + 1500 + 2000 + 2500 + 3000 = 10000


### 2. terbesar.go
```go
package main

import "fmt"

func main() {
	var n, bilangan, bilanganTerbesar int

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&bilangan)

		if i == 1 || bilangan > bilanganTerbesar {
			bilanganTerbesar = bilangan
		}
	}

	fmt.Print(bilanganTerbesar)
}
```

#### Deskripsi
`terbesar.go` Merupakan program untuk mencari bilangan terbesar dari `n`buah bilangan bulat 
```go
	fmt.Scan(&n)
```
Masukkan `n`
```go
for i := 1; i <= n; i++ {
    fmt.Scan(&bilangan)

    if i == 1 || bilangan > bilanganTerbesar {
        bilanganTerbesar = bilangan
    }
}
```
lalu input `bilangan` dimasukkan ke dalam perulangan supaya bisa menginputkan `bilangan` sejumlah `n` kali.

Terdapat kondisi 
```go
if i == 1 || bilangan > bilanganTerbesar {
    bilanganTerbesar = bilangan
}
```
pada iterasi pertama atau ketika `bilangan > bilanganTerbesar`: `bilanganTerbesar` diperbarui dengan nilai `bilangan` yang terbaru.

`i == 1` ini penting untuk mengatasi jika masukkan `bilangan` nya negatif semua karena kondisi `bilangan > bilanganTerbesar` tidak akan terpenuhi sehingga nilai `bilanganTerbesar` nya akan tetap `0`.

#### Contoh Input Output
#### Input
```text 
5
5 9 2 7 5
```
#### Output
```text
9
```

### 3. bioskop.go
```go
package main

import "fmt"

func main() {
	var (
		jumlahTiket, jumlahTransaksi, totalPendapatan, harga, diskon int
		jenis, hari                                                  string
	)

	fmt.Scan(&jumlahTransaksi)
	for i := 1; i <= jumlahTransaksi; i++ {
		harga = 0
		fmt.Printf("Transaksi ke-%d\n", i)
		fmt.Print("Masukkan jenis film (Reguler, 3D, IMAX): ")
		fmt.Scan(&jenis)

		fmt.Print("Hari: ")
		fmt.Scan(&hari)

		fmt.Print("Jumlah Tiket: ")
		fmt.Scan(&jumlahTiket)

		if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
			if jenis == "Reguler" {
				harga = 40000
			} else if jenis == "3D" {
				harga = 60000
			} else if jenis == "IMAX" {
				harga = 80000
			} else {
				fmt.Println("Jenis film tidak valid")
			}
		} else if hari == "Jumat" || hari == "Sabtu" || hari == "Minggu" {
			if jenis == "Reguler" {
				harga = 50000
			} else if jenis == "3D" {
				harga = 70000
			} else if jenis == "IMAX" {
				harga = 90000
			} else {
				fmt.Println("Jenis film tidak valid")
			}
		} else {
			fmt.Println("Hari tidak valid")
		}

		totalHarga := harga * jumlahTiket

		// Diskon ketika beli tiket 5 keatas
		if jumlahTiket >= 5 {
			diskon = totalHarga * 10 / 100
			totalHarga = totalHarga - diskon
		}

		totalPendapatan = totalPendapatan + totalHarga

		fmt.Printf("\n====================\n")
		fmt.Println("Total Harga:",totalHarga)
	}

	fmt.Println("Total Pendapatan:",totalPendapatan)
	fmt.Println("====================")
}
```

#### Deskripsi
`biskop.go` Merupakan program untuk menghitung total pendapatan bioskop dari beberapa transaksi pembelian tiket dalam satu periode. Program ini akan menghitung total transaksi berdasarkan harga per tiketnya yang berbeda-beda tergantung pada `jenis film`, dan `hari` nya. Jika pembelian dalam satu transaksi >= 5 tiket, pembeli akan mendapatkan diskon `10%`.

Masukkan jumlah transaksi juga sebagai batas iterasi pada perulangan.
```go
fmt.Scan(&jumlahTransaksi)
```

Masukkan `jenis film` , hari, dan `jumlah tiket` pada masing-masing transaksi nya.
```go
for i := 1; i <= jumlahTransaksi; i++ {
    fmt.Printf("Transaksi ke-%d\n", i)
    fmt.Println("Masukkan jenis film (Reguler, 3D, IMAX): ")
    fmt.Scan(&jenis)

    fmt.Println("Hari: ")
    fmt.Scan(&hari)

    fmt.Println("Jumlah Tiket")
    fmt.Scan(&jumlahTiket)

    ...
}
```

Pengkondisian untuk menentukan harga tiket sesuai dengan `hari pemutaran`, dan `jenis film` yang dimasukkan sebelumnya.

Hari Senin - Kamis 
- harga 1 tiket pada jenis `Reguler` = 40.000
- harga 1 tiket pada jenis `3D` = 60.000
- harga 1 tiket pada jenis `IMAX` = 80.000

Hari Jumat - Minggu
- harga 1 tiket pada jenis `Reguler` = 50.000
- harga 1 tiket pada jenis `3D` = 70.000
- harga 1 tiket pada jenis `IMAX` = 90.000
```go
if hari == "Senin" || hari == "Selasa" ||hari == "Rabu" || hari == "Kamis" {
    if jenis == "Reguler" {
        harga = 40000
    } else if jenis == "3D" {
        harga = 60000
    } else if jenis == "IMAX" {
        harga = 80000
    }
}

if hari == "Jumat" || hari == "Sabtu" || hari == "Minggu" {
    if jenis == "Reguler" {
        harga = 50000
    } else if jenis == "3D" {
        harga = 70000
    } else if jenis == "IMAX" {
        harga = 90000
    }
}

totalHarga := harga * jumlahTiket // Hitung total harga
```

Jika tiket yang dibeli >= 5, maka dapat diskon 10% dan `totalHarga` dihitung kembali 
```go
// Diskon ketika beli tiket 5 keatas
if jumlahTiket >= 5 {
    diskon = totalHarga * 10 / 100
    totalHarga = totalHarga - diskon
}
```
Lalu totalkan pendapatan dari setiap transaksi nya
```go
totalPendapatan = totalPendapatan + totalHarga
```

Total harga dari masing-masing transaksi ditampilkan
```go
fmt.Printf("\n====================\n")
fmt.Println("Total Harga:",totalHarga)
```
dan total pendapatan ditampilkan diluar perulangan supaya hanya ditampilkan 1 kali
```go
fmt.Println("Total Pendapatan:",totalPendapatan)
fmt.Println("====================")
```

#### Contoh Input, Output.
#### Input
```text
jumlahTransaksi: 1

Reguler
Senin
2
```

#### Output
```text
====================
Total Harga: 80000
Total Pendapatan: 80000
====================
```
Hari = Senin, dan jenis = Reguler -> harga = 40.000/tiket
40.000 x 2 (jumlah tiket) = 80.000

## Unguided
### 1. prima.go
```go
package main

import "fmt"

func main() {
	var (
		n, jumlahFaktor int
	)
	fmt.Println("Masukkan Bilangan:")
	fmt.Scan(&n)

	if n < 1 {
		fmt.Println("Masukkan minimal angka 1")
		return
	}

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Println(i)
			jumlahFaktor++
		}

		if jumlahFaktor == 2 {
			fmt.Printf("\n%d Merupakan bilangan prima", n)
		} else {
			fmt.Printf("\n%d Bukan bilangan prima", n)
		}
	}
}
```

#### Output
![Screenshot Output Unguided](/praktikum/05-perulangan-for-loop/unguided/prima/output.png)

#### Deskripsi
`prima.go` Merupakan program sederhana untuk menampilkan faktor dari `bilangan` yang dimasukkan dan untuk menentukan bilangan tersebut merupakan bilangan prima atau bukan. 

- Jika iya: akan menampilkan "`n` Merupakan bilangan prima"
- Jika bukan: akan menampilkan "`m` Bukan bilangan prima"

Masukkan bilangan mulai dari angka 1
```go
fmt.Println("Masukkan Bilangan:")
fmt.Scan(&n)

if n < 1 {
    fmt.Println("Masukkan minimal angka 1")
    return
}
```
Perulangan dan pengkondisian didalamnya untuk menentukan bilangan `faktor` dari bilangan `n` dan apakah `n` bilangan prima atau bukan.
```go
for i := 1; i <= n; i++ {
    if n%i == 0 {
        fmt.Println(i)
        jumlahFaktor++
    }

    if jumlahFaktor == 2 {
        fmt.Printf("\n%d Merupakan bilangan prima", n)
    } else {
        fmt.Printf("\n%d Bukan bilangan prima", n)
    }
}
```
Pengkondisian untuk mencari `faktor` dari bilangan `n`
```go
if n%i == 0 {
    fmt.Println(i)
    jumlahFaktor++
}
```
Penjelasan: Jika `n` habis dibagi dengan iterasi perulangan saat ini (i), `i` tersebut adalah faktor dari `n` lalu `jumlahFaktor` ditambah 1 nilai.

Pengkondisian untuk menentukan apakah `n` adalah bilangan prima atau bukan
```go
if jumlahFaktor == 2 {
    fmt.Printf("\n%d Merupakan bilangan prima", n)
} else {
    fmt.Printf("\n%d Bukan bilangan prima", n)
}
```
Penjelasan: Bilangan Prima adalah bilangan yang hanya akan habis dibagi dengan 1 dan dengan dirinya sendiri. Artinya, bilangan prima hanya punya 2 faktor. Jika jumlahFaktor yang nilainya ditambahkan pada pengkondisian sebelumnya hanya ada 2 maka bilangan `n` adalah `bilangan prima`. 

### 2. ecct.go
```go
package main

import "fmt"

// Function untuk range nilai 1-4
func bacaNilai(label string) int {
	var nilai int

	for {
		fmt.Printf("%s: ", label)
		fmt.Scan(&nilai)

		// Rentang nilai 1-4
		if nilai >= 1 && nilai <= 4 {
			return nilai
		}

		fmt.Println("Masukkan nilai pada rentang 1-4. Coba lagi")
	}
}

func main() {
	var (
		i, m, p, v, ga, fs, ci, tc, mahasiswaLulus int
		avg float64
		hasil []string
	)

	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&m)


	for i = 1; i <= m; i++ {
		fmt.Printf("Masukkan nilai mahasiswa %d (1-4)\n", i)

		// panggil func bacaNilai untuk nilai supaya rentang 1-4 saja
		p = bacaNilai("P")
		v = bacaNilai("V")
		ga = bacaNilai("GA")
		fs = bacaNilai("FS")
		ci = bacaNilai("CI")
		tc = bacaNilai("TC")

		// total dan avg
		totalNilai := p + v + ga + fs + ci + tc
		avg = float64(totalNilai) / 6

		fmt.Printf("Nilai rata-rata mahasiswa ke-%d: %.2f\n\n", i, avg)
		if avg >= 3.33 {
			hasil = append(hasil, "Mendapatkan Sertifikat")
			mahasiswaLulus++
		} else {
			hasil = append(hasil, "Tidak Mendapatkan Sertifikat")
		}
	}

	for _, sertifikat := range hasil {
		fmt.Println(sertifikat)
	}

	fmt.Printf("%d Mahasiswa mendapatkan sertifikat", mahasiswaLulus)
}
```

#### Output
![Screenshot Output Unguided](/praktikum/05-perulangan-for-loop/unguided/ecct/output.png)


#### Deskripsi
`ecct.go` Merupakan program untuk menentukan apakah mahasiswa mendapatkan sertifikat ECCT atau tidak, dengan minimal nilai rata-rata: `3.33` yang  dihitung dari 6 nilai berikut ini:
1. P
2. V
3. GA
4. FS
5. CI
6. TC
dengan rentang penilaian 1-4. Program ini juga menampilkan ada berapa mahasiswa yang mendapatkan sertifikat ECCT. 

Masukkan jumlah mahasiswa. Lalu memasukkan 6 nilai yang telah disebutkan pada masing-masing mahasiswa di tiap iterasi perulangan.
```go
fmt.Print("Masukkan jumlah mahasiswa: ")
fmt.Scan(&m)

for i = 1; i <= m; i++ {
    fmt.Printf("Masukkan nilai mahasiswa %d (1-4)\n", i)
    p = bacaNilai("P")
    v = bacaNilai("V")
    ga = bacaNilai("GA")
    fs = bacaNilai("FS")
    ci = bacaNilai("CI")
    tc = bacaNilai("TC")
}
```
Disini saya menggunakan pendekatan untuk membuat fungsi lagi di luar `func main()` untuk membatasi rentang nilai 1-4 supaya tidak ribet mengimplementasikan kondisinya pada masing-masing variabel. 

Fungsi `bacaNilai` digunakan untuk menginputkan nilai yang value nya akan disimpan ke variabel `nilai` pada function `bacaNilai` dan disimpan ke variabel nya masing-masing (P, V, GA, FS, CI, TC).

Fungsi `bacaNilai` juga membatasi nilai yang dimasukkan berada pada rentang 1-4.
```go
if nilai >= 1 && nilai <= 4 {
    return nilai
}

fmt.Println("Masukkan nilai pada rentang 1-4. Coba lagi")
```

Keenam nilai yang dimasukkan sebelumnya ditotalkan dan dihitung nilai rata-ratanya. Jika >= 3.33 maka mahasiswa tersebut dinyatakan `Mendapatkan Sertifikat` jika tidak `Tidak Mendapatkan Sertifikat` dan jumlah mahasiswa yang lulus ditambah 1 supaya diakhir bisa ditotalkan
```go
totalNilai := p + v + ga + fs + ci + tc
avg = float64(totalNilai) / 6

if avg >= 3.33 {
    hasil = append(hasil, "Mendapatkan Sertifikat")
    mahasiswaLulus++
} else {
    hasil = append(hasil, "Tidak Mendapatkan Sertifikat")
}
```

Lalu menampilkan hasil nya (Mendapatkan sertifikat atau tidak) dan berapa banyak mahasiswa yang berhasil mendaparkan sertifikat
```go
for _, sertifikat := range hasil {
    fmt.Println(sertifikat)
}

fmt.Printf("%d Mahasiswa mendapatkan sertifikat", mahasiswaLulus)
```
## Kesimpulan
Dari Praktikum minggu ke-5 ini, saya berhasil memahami cara kerja dari beberapa bentuk perulangan `for` pada bahasa Go, dan dapat mengimplementasikan nya secara langsung. Di Go, hanya memiliki satu kata kunci perulangan yaitu `for`. Go tidak mempunyai `while`, dan `do-while` seperti pada beberapa bahasa pemrograman lain. Namun, kita tetap bisa menggunakan pendekatan yang sama dengan `while` dan `do-while` walaupun tetap menggunakan kata kunci `for` dan dengan cara penulisan yang berbeda pada perulangan nya sesuai dengan pendekatannya masing-masing seperti:
- `for` dengan 3 komponen
- `for` hanya dengan kondisi 
- `for` tanpa menggunakan kondisi
- dan `for-range` untuk suatu kumpulan data seperti array,slice,dan map

## Referensi
#### 1. Modul 05 Praktikum - Perulangan For