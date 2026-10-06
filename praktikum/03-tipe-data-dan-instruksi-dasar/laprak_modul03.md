# <h1 align="center">Laporan Praktikum Modul 03 - Variabel dan Operator</h1>
<p align="center">Muhamad Nawaf Abduh - 109092630003</p>

## Dasar Teori

### 1. Variabel
Variabel merupakan wadah untuk menyimpan data yang nilainya dapat digunakan atau diubah selama program berjalan. 

#### Deklarasi Variabel
Deklarasi variabel di Go, dapat menggunakan kata kunci `var` lalu `namaVariabel` nya, lalu `tipeData` nya. Cara penulisan:

```go
var namaVariabel tipeData
```
lalu kita isi `variabel` tersebut sehingga dia memiliki nilai nya. Contoh:

```go
var name string
name = "Muhamad Nawaf Abduh"
```

di Go, kita juga dapat langsung memberikan nilai pada `variabel` nya. Contoh:
```go
var angka int = 10
```

di Go, kita juga dapat mendeklarasikan beberapa `variabel` sekaligus.
```go
var a, b, c int
```

#### Deklarasi Variabel dengan Singkat
Di Go, variabel juga dapat dideklarasikan dengan bentuk singkat menggunakan `:=`. Pada bentuk ini, tipe data ditentukan secara otomatis oleh Go berdasarkan nilai yang diberikan di dalam variabel nya, dan kita tidak perlu menuliskan kata kunci `var` untuk mendeklarasikan variabel nya..

```go
angka := 10
pesan := "Halo"
```
Dengan menggunakan `:=` kita langsung mendeklarasikan `variabel` sekaligus menginisialisasikan datanya secara bersamaan. Dan Go secara otomatis dapat mengetahui tipe data apa yang digunakan pada `variabel` yang kita deklarasikan berdasarkan nilai di dalam `variabel` nya. 

Catatan: `:=` hanya bisa digunakan didalam function lokal. `:=` tidak bisa digunakan di luar function (tidak dapat digunakan sebagai `variabel global`)

#### Mengakses Variabel
Setelah variabel dideklarasikan di Go, variabel tersebut benar-benar harus kita gunakan. Tidak boleh tidak. Kita dapat mengakses nilai di dalamnya dengan menuliskan nama variabel tersebut. Contohnya, variabel `angka` dan `pesan` dapat digunakan sebagai berikut:

```go
var angka int = 10
pesan := "Halo"

fmt.Println(angka)
fmt.Println(pesan)
```

`fmt.Println(angka)` digunakan untuk menampilkan nilai yang tersimpan di dalam variabel `angka`.

### 2. Operator
Operator merupakan simbol yang digunakan untuk melakukan operasi terhadap satu atau beberapa nilai atau variabel. ada beberapa jenis Operator, yaitu:

#### Operator Aritmatika
Operator aritmatika digunakan untuk melakukan perhitungan matematika dasar. Operator aritmatika terdiri dari :
`+, -, *, /, %`
- `+` untuk penjumlahan
- `-` untuk pengurangan
- `*` untuk perkalian
- `/` untuk pembagian bilangan bulat
- `%` untuk menghitung sisa bagi
Contoh penggunaan:
```go
a := 10
b := 3

fmt.Println(a + b) // = 13
fmt.Println(a - b) // = 7
fmt.Println(a * b) // = 30
fmt.Println(a / b) // = 3. hasilnya 3 karena 3 * 3 lah yang paling mendekati 10
fmt.Println(a % b) // = 1
```
Operator aritmatika ini dapat digabungkan dengan `=` Contoh:

```go
var a, b int

a = 10
b = 2

a += b
a -= b
a *= b
a /= b
a %= b
```
Pendekatan ini dapat digunakan untk mempersingkat penulisan operasi terhadap `variabel` ketika mmebuat sebuah program Go. `a += b` memiliki arti yang sama dengan `a = a + b` begitu pun yang lainnya.

#### Operator Perbandingan
Operator perbandingan digunakan untuk membandingkan dua nilai. Hasil dari perbandingan akan berupa nilai `bool`, yaitu `true` atau `false`. Operator perbandingan terdiri dari: `==, !=, <, >, <=, >=`
- `==` sama dengan
- `!=` tidak sama dengan
- `<` lebih kecil dari
- `>` lebih besar dari
- `<=` lebih kecil sama dengan
- `>=` lebih besar dari sama dengan

Contoh:
```go
a := 10
b := 3

fmt.Println(a == b) // false
fmt.Println(a != b) // true
fmt.Println(a < b)  // false
fmt.Println(a > b)  // true
fmt.Println(a <= b) // true
fmt.Println(a >= b) // false
```

#### Operator Logika
Operator logika merupakan operator yang biasa digunakan untuk menghubungkan atau membalikkan suatu kondisi. Operator logika terdiri dari:
- `&&`: `AND` 
- `||`: `OR`
- `!`: `NOT`

#### &&
Akan menghasilkan `true` jika kedua kondisi bernilai `true`

#### ||
Akan menghasilkan `true` ketika setidaknya salah satu kondisi bernilai `true`

#### !
Digunakan untuk membalikkan nilai dari kondisi yang awalnya `true` menjadi `false`, begitupun sebaliknya

Contoh penggunaan:
```go
a := 10
b := 3

fmt.Println(a > 5 && b < 5)  // hasilnya true
fmt.Println(a < 5 || b < 5)  // hasilnya true
fmt.Println(!(a > 5))        // hasilnya false
```

#### Operator Bitwise
Operator bitwise digunakan untuk melakukan operasi terhadap bit-bit pada bilangan bulat. Operator bitwise terdiri dari:
- `&`: `AND` bitwise
- `|`: `OR` bitwise
- `^`: `XOR` bitwise atau membalik bit jika digunakan sebagai operator unary
- `&^`: `AND NOT` bitwise
- `<<`: menggeser bit ke kiri
- `>>`: menggeser bit ke kanan

Contoh penggunaan:
```go
a := 6 // biner: 0110
b := 3 // biner: 0011

fmt.Println(a & b)  // = 2 (0010)
fmt.Println(a | b)  // = 7 (0111)
fmt.Println(a ^ b)  // = 5 (0101)
fmt.Println(a &^ b) // = 4 (0100)
fmt.Println(a << 1) // = 12 (1100)
fmt.Println(a >> 1) // = 3 (0011)
```

## Guided
### 1. konversi.go

```go
package main

import "fmt"

func main() {
	var (
		c, k float64
	)

	// Input
	fmt.Println("Masukkan suhu celcius")
	fmt.Scanln(&c)

	// Konversi suhu celcius ke kelvin
	k = c + 273

	// Output
	fmt.Print(k)
}
```
#### Deskripsi
`konversi.go` merupakan program yang dibuat untuk mengkonversi suhu dari satuan `Celcius` ke satuan `Kelvin`. Kita menginputkan suhu dalam satuan `celcius` 
```go
	fmt.Scanln(&c)
```
lalu, nilai `celcius` yang kita inputkan, dihitung menggunakan rumus `k = c + 273` untuk dikonversikan kedalam satuan `kelvin`
```go
	k = c + 273
``` 
dan nilai `celcius` yang telah dikonversikan ke dalam suhu `kelvin` ditampilkan menggunakan
```go
	fmt.Print(k)
```
#### Contoh Input, Output
#### Input
```text
31
```
#### Output
```text
304
```

### 2. tukar.go
```go
package main

import "fmt"

func main() {
	var (x,y,z int)

	// Input nilai a dan b
	fmt.Print("Masukkan nilai x: ")
	fmt.Scanln(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scanln(&y)
	fmt.Print("Masukkan nilai z: ")
	fmt.Scanln(&z)

	// Tukar nilai x,y,z
	temp := x
	x = z
	z = y
	y = temp

	// Output hasil setelah ditukar
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
```

#### Deskripsi
`tukar.go` merupakan program untuk menukar-nukarkan nilai bilangan bulat pada variabel `x`, `y`, dan `z`.

Pertukaran dilakukan dengan cara:

1. `temp := x` menyimpan nilai awal `x` ke dalam variabel sementara `temp`, agar nilai `x` diawal tidak hilang saat `x` diubah.
2. `x = z` `x` diisi dengan nilai `z`.
3. `z = y` `z` diisi dengan nilai `y`.
4. `y = temp` `y` diisi dengan nilai awal `x` yang tadi disimpan di variabel `temp`.

Maka, nilai akhirnya adalah:
- `x` berisi nilai awal `z`, 
- `y` berisi nilai awal `x`, 
- `z` berisi nilai awal `y`.

#### Contoh Input, Output.
#### Input
```text
x = 3
y = 1
z = 2
```

#### Output
```text
x = 2
y = 3
z = 1
```

### 3. kasir.go
```go
package main

import "fmt"

func main () {
	var (x, lembar10rb, lembar5rb, lembar1rb int)

	// Input
	fmt.Println("Masukkan nominal uang: ")
	fmt.Scanln(&x)

	// Operasi
	lembar10rb = x / 10000 
	x %= 10000 
	lembar5rb = x / 5000
	x %= 5000
	lembar1rb = x / 1000

	// Output
	fmt.Println(lembar10rb)
	fmt.Println(lembar5rb)
	fmt.Println(lembar1rb)
}
```
#### Deskripsi
`kasir.go` merupakan program untuk menghitung jumlah lembar uang `10000`, `5000`, dan `1000` dari nominal yang dimasukkan. Atau, pada kode program saya, dideklarasikan sebagai variabel `x`.

#### Contoh Input Output
#### Input
```text
19000
```
`x = 19000`

#### Output
```text
1, 1, 4
```
Artinya, 1 lembar `10000`, 1 lembar `5000`, dan 4 lembar `1000`. Jika dijumlah maka hasilnya `19000`. Sesuai dengan yang diinputkan

## Unguided

### 1. konversiSuhu.go
```go
package main
import "fmt"

func main () {
	var celcius,raemur float64

	fmt.Println("Konversi suhu Celcius ke Raemur. Masukkan suhu Celcius: ")
	fmt.Scanln(&celcius)

	raemur = 4.0/5.0 * celcius

	fmt.Printf("%.2f\n", raemur)
}
```

#### Output
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/unguided/konversi-suhu/output.png)

#### Deskripsi
`konversiSuhu.go` merupakan program untuk mengkonversi nilai suhu yang diinputkan dalam satuan `celcius` ke satuan `raemur` yang dihitung menggunakan rumus `4/5 * celcius`
```go 
raemur = 4.0/5.0 * celcius
```
Hasil perhitungan disimpan dalam variabel `raemur`, lalu ditampilkan menggunakan `fmt.Printf("%.2f\n", raemur)`. Format `%.2f` membatasi hanya dua angka di belakang koma.
```go
fmt.Printf("%.2f\n", raemur)
```

#### Contoh Input, Output
#### Input
Menginputkan nilai `celcius`
```text
100
```
`celcius = 100`

#### Output
```text
80.00
```
`nilai raemur`

### 2. konversiHari.go

```go
package main
import "fmt"

func main () {
	var (tahun,bulan,minggu,hari int)

	// Input hari
	fmt.Println("Masukkan jumlah hari: ")
	fmt.Scan(&hari)

	tahun = hari / 360
	hari %= 360
	bulan = hari / 30
	hari %= 30
	minggu = hari / 7
	hari %= 7

	fmt.Printf("\n %d Tahun, %d Bulan, %d Minggu, %d Hari", tahun, bulan, minggu, hari)
}
```

#### Output
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/unguided/konversi-hari/output.png)


#### Deskripsi
`konversiHari.go` merupakan program untuk menghitung jumlah tahun,bulan,minggu dan hari dari `hari` yang diinputkan. Semua itu dihitung menggunakan perhitungan :
- 1 tahun = 12 bulan atau 360 hari
- 1 bulan = 30 hari
- 1 minggu = 7 hari

#### Contoh Input, Output
#### Input
```text
400
```
`hari = 400`

#### Output
```text
1 Tahun, 1 Bulan, 1 Minggu, 3 Hari
```

## Kesimpulan

Berdasarkan hasil praktikum dan pembahasan mengenai Variabel dan Operator pada modul 3, dapat disimpulkan bahwa:

### 1. Variabel 
Variabel merupakan wadah untuk menyimpan data yang nilainya dapat digunakan atau diubah selama program berjalan.o Di G, variabel dapat dideklarasikan menggunakan kata kunci `var` dengan menuliskan nama dan tipe datanya. Kita juga dapat menggunakan bentuk singkat `:=` untuk mendeklarasikan variabel sekaligus memberikan nilai, dan tipe datanya ditentukan secara otomatis oleh Go berdasarkan nilai didalam variabel nya.

### 2. Operator
Operator merupakan simbol yang digunakan untuk melakukan operasi terhadap satu atau beberapa `nilai` atau `variabel`. Ada beberapa operator :

- `Operator aritmatika` digunakan untuk melakukan perhitungan matematika dasar
- `Operator perbandingan` digunakan untuk membandingkan dua nilai yang akan menghasilkan nilai `boolean`, yaitu `true` atau `false` 
- `Operator logika` digunakan untuk menghubungkan dua kondisi atau membalikkan suatu kondisi yang akan menghasilkan nilai `boolean` berdasarkan tabel kebenarannya
- `Operator bitwise` digunakan untuk melakukan operasi terhadap bit-bit pada bilangan bulat

Dari praktikum ini, saya berhasil memahami keseluruhan materi dengan baik, karena saya bisa sambil praktik langsung, sehingga lebih mudah untuk memahaminya

## Referensi
#### 1. Materi Pekan 2 di LMS - Skema dan Struktur Algoritma