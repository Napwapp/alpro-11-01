# <h1 align="center">Tugas Pendahuluan Modul 05 - Perulangan For</h1>
<p align="center">Muhamad Nawaf Abduh - 109092630003</p>

### 1. Apa fungsi utama dari struktur perulangan for dalam program Go? Sebutkan dan
jelaskan tiga komponen pada perulangan for (inisialisasi, kondisi, dan post statement)!

#### Jawab: 
`for` dalam `Go` merupakan satu satunya cara untuk mengeksekusi kode secara berulang (perulangan).

```go
for ... {
    ...
}
```

Kode yang berada didalam `for{}` akan dieksekusi secara terus menerus dan berulang hingga suatu kondisi yang kita tentukan terpenuhi.

3 Komponen Perulangan:
- Inisialisasi
- Kondisi
- Post Statement

1. Inisialisasi
Mendeklarasikan suatu variabel dan menginisialisasikan nilai awalnya supaya dapat dijadikan sebagai penghitung pada kondisi perulangannya.

```go
for i := 0; ... {
    ...
}
```
Varibale `i` dideklarasikan dan menginisialiasi `0` sebagai nilai awalnya

2. Kondisi
Kondisi pada perulangan digunakan untuk menentukan kapan perulangan akan dieksekusi dan kapan akan berhenti. Contoh:
`i <= 5`

```go
for i := 0; i <= 5;... {
    ...
}
```
Sebelumnya, variabel `i` kita menginisialisasikan `0` sebagai nilai awalnya
`i = 0`

kondisinya  `i <= 5` maka perulangan diatas akan dijalankan selama nilai `i` nya tidak lebih dari `5`

Perulangan akan selalu dijalankan kembali selama kondisi nya masih bernilai `true`.

3. Post Statement
Post statement merupakan pernyataan yang digunakan untuk menaikkan dan menurunkan nilai variabel ny
```go
for i := 0; i <= 5; i++ {
    ...
}
```
- `++` akan menambahkan 1 nilai ke variabel `i` 
- `++` akan mengurangkan 1 nilai ke variabel `i` 

### 2. Bahasa Go tidak memiliki perulangan while dan do-while. Jelaskan perbedaan antara for dengan tiga komponen dan for yang hanya menggunakan kondisi! Berikan contoh kode masing-masing.

#### Jawab: 

- for dengan menggunakan 3 komponen
Jika menggunakan for dengan tiga komponen, kita sudah mengetahui dan langsung membatasi kapan perulangan tersebut akan dihentikan. Contoh:
```go
for i := i; i <= 5; i++ {
    fmt.Println(i)
}
```

- for yang hanya menggunakan kondisi
```go
i = 1 
for i <= 5 {
    fmt.Println(i)
    i++
} 
```
Contoh perulangan diatas sebenarnya sama saja dengan `for` yang menggunakan 3 komponen, namun hanya dibedakan dengan cara penulisannya saja. Perulangan diatas, hanya menggunakan kondisi saja, yaitu `i <= 5`. Artinya, perulangan akan terus dieksekusi selama kondisi tersebut bernilai `true`

### 3. ganjil.go
```go
package main

import "fmt"

func main() {
	var n, total int
	fmt.Println("Masukkan bilangan:")
	fmt.Scan(&n)

	for i := 1; i <= n; i += 2 {
		fmt.Printf("%d ", i)
		total = total + i
	}
	fmt.Printf("\nTotal: ")
	fmt.Printf("%d\n", total)
}
```

#### Output
![Screenshot Output Unguided](/praktikum/05-perulangan-for-loop/tp/ganjil/output.png)

#### Deskripsi
`ganjil.go` merupakan program untuk menampilkan bilangan ganjil mulai dari 1 hingga ke `n` yang dimasukkan. Lalu semua bilangan ganjil tersebut di jumlahkan. 

1. Masukkan bilangan `n`
`fmt.Scan(&n)`

2. 
```go
for i := 1; i <= n; i += 2 {
	fmt.Printf("%d ", i)
	total = total + i
}
```
Menggunakan perulangan `for`.

- `i := 1;` Mendeklarasikan sekaligus menginisialisasi nilai awal variabel `i` 
- `i <= n;` Kondisi perulangan. Selama `i` tidak lebih dari `n` perulangan akan terus dieksekusi
- `i += 2` nilai `i` ditambah dengan 2. Supaya nilai `i` hanya bilangan `ganjil`
- `fmt.Printf("%d ", i)` dan `total = total + i`. Menampilkan nilai dari variabel `i` lalu menghitung total dengan menjumlahkan variabel `total` dengan `i` 
- `fmt.Printf("%d\n", total)` Menampilkan nilai total

### 4. Perkalian.go
```go
package main
import "fmt"

func main () {
	var a,b int

	fmt.Println("Masukkan bilangan a untuk dikalikan sampai ke bilangan b:")
	fmt.Scan(&a, &b)
	
	for i := 1; i <= b; i++ {
		hasilPerkalian := a * i
		fmt.Printf("%d x %d = %d\n", a, i, hasilPerkalian)
	}
}
```

#### Output
![Screenshot Output Unguided](/praktikum/05-perulangan-for-loop/tp/perkalian/output.png)

#### Deskripsi
`perkalian.go` Merupakan program perkalian yang akan mengalikan bilangan `a` dengan 1 hingga ke bilangan `b`

1. `fmt.Scan(&a, &b)` input bilangan a dan b
2. Menggunakan perulangan `for`
```go
for i := 1; i <= b; i++ {
	hasilPerkalian := a * i
	fmt.Printf("%d x %d = %d\n", a, i, hasilPerkalian)
}
```
- `i := 1;` Mendeklarasikan sekaligus menginisialisasi nilai awal variabel `i` 
- `i <= n;` Kondisi perulangan. Selama `i` tidak lebih dari `b` perulangan akan terus dieksekusi
- `i++` nilai `i` ditambah 1
- hasil perkalian `a * i` disimpan divariabel `hasilPerkalian`. `perkalian := a * i`
- Menampilkan output dengan format `a x i = hasilPerkalian`

### Kesimpulan
Perulangan merupakan suatu struktur kontrol yang akan mengeksekusi kode didalamnya selama kondisi perulangannya bernilai `true`. Perulangan di Go ditulis menggunakan keywords `for`. Dan di Go, tidak memiliki perulangan `while`, dan `do-while` seperti pada pemrograman bahasa lain. Namun, di Go bisa memakai konsep yang sama dengan `while` dan `do-while` dengan tetap menggunakan `for` namun dengan cara penulisan yang berbeda.