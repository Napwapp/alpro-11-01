# <h1 align="center">Tugas Pendahuluan Modul 04 - Runtunan / Sekuensi</h1>
<p align="center">Muhamad Nawaf Abduh - 109092630003</p>

### Soal 1 – Evaluasi Ekspresi Kontrol dalam Go
`soal1.go`
```go
package main
import "fmt"

func main () {
	intNum := 5
	intOther := 10
	var sngNum float64 = -3
	
	if intOther > 0 && intNum > 0 || sngNum > 0 {
		fmt.Println("Beep")
	} 
}
```

##### Output
![Screenshot output sisa kue](/praktikum/04-runtunan-sekuensi/tp/soal1/output.png)


#### Deskripsi
Soal 1 – Evaluasi Ekspresi Kontrol dalam Go

Diberikan segmen kode berikut, di mana setiap ekspresi kontrol yang terdaftar di bawah
dapat digunakan sebagai kondisi pada pernyataan if:

```go
intNum := 5
intOther := 10
var sngNum float64 = -3
if ... {
 fmt.Println("Beep")
}
```

Soal No.1 mengharuskan kita untuk menentukan apakah masing-masing ekspresi kontrol tersebut menghasilkan
`true` atau `false`. Tuliskan hasilnya dengan kata `true` atau `false`

1. `intNum > 5` = **false**
2. `intNum >= 5 && intOther < 11` = **true**
3. `sngNum != -1 || intOther < 0` = **true**
4. `!(intNum > 3) || intNum <= 5` = **true**
5. `!(intOther >= intNum)` = **false**
6. `0 - sngNum > 0` = **true**
7. `4 / 2 == intOther / intNum` = **true**
8. `intOther % 2 == 0` = **true**
9. `intOther + 2 * intNum != 30 || !(sngNum > 0)` = **true**
10. `intOther > 0 && intNum > 0 || sngNum > 0` = **true**
11. `sngNum > 0 || (intNum >= 0 && -1 * intOther == -10)` = **true**
12. `intNum == 5` = **true**
13. `intNum > 0 || (sngNum <= 0 && intOther == 13)` = **true**
14. `!(!(!(!(intNum > 0))))` = **true**

#### Kondisi yang saya pilih untuk membuat program nya
`intOther > 0 && intNum > 0 || sngNum > 0` atau nomor `10`

```go
package main
import "fmt"

func main () {
	intNum := 5
	intOther := 10
	var sngNum float64 = -3
	
	if intOther > 0 && intNum > 0 || sngNum > 0 {
		fmt.Println("Beep")
	} 
}
```
Kondisi tersebut bernilai `true`, maka `Beep` akan ditampilkan.
```go
fmt.Println("Beep")
```

### Soal 2 – Tracing: Evaluasi Pernyataan Kondisi
`soal2.go`

```go
// Soal 2 – Tracing: Evaluasi Pernyataan Kondisi
package main

import "fmt"

func main() {
	x := 10
	y := 5
	z := 15
	result := 0

	if x > 5 { // true
		if y < 10 { // true
			result = x + y // = 15
		} else {
			result = x - y // false
		}
	}

	if z > 10 && x == 10 { // true
		result += z // = 30
	} else { // false
		result = z - x 
	}

	if x == 10 || y > 10 {
		result += 5 // = 35
	} else if y == 5 && z > 10 { 
		result -= 5 // false
	} else {
		result *= 2
	}

	if !(x < 15 && y < 10) {
		result += 10 // false
	} else {
		result -= 10 // 25
	}

	fmt.Println("Nilai akhir result:", result)
}
```

##### Output
![Screenshot output boolean](/praktikum/04-runtunan-sekuensi/tp/soal2/output.png)


#### Deskripsi
Di Soal 2 ini, saya harus untuk menganalisis alur eksekusi program dan menentukan nilai akhir dari setiap variabel setelah program selesai dijalankan. Dan menuliskan juga output yang dihasilkan oleh program. Jawaban akhirnya adalah:
- Nilai dari variabel `x`,`y`, dan `z` tetap sama hingga program selesai dijalankan
- Nilai dari variabel `result` terus berubah-rubah sesuai pengkondisian yang dieksekusi. Dan nilai akhir dari variabel `result` adalah `25`.

Catatan:
```text
Untuk alur eksekusi program yang lebih detailnya ada pada : praktikum\04-runtunan-sekuensi\tp\soal2\soal2.md
```

### Soal 3 – Menentukan Jumlah Hari dalam Sebulan Berdasarkan Tahun dan Bulan
`soal3.go`

```go
package main

import "fmt"

func main() {
	var (
		tahun, hari  int
		bulan string
	)

	fmt.Print("Masukkan tahun: ")
	fmt.Scanln(&tahun)
	fmt.Print("Masukkan bulan dalam format (Jan, Feb, Mar, Apr, Mei, Jun, Jul, Agu, Sep, Okt, Nov, Des): ")
	fmt.Scanln(&bulan)

	switch bulan {
		case "Jan":
			hari = 31
		case "Feb":
			if (tahun%4 == 0 && tahun%100 != 0) || (tahun%400 == 0) {
				hari = 29
			} else {
				hari = 28
			}
		case "Mar":
			hari = 31
		case "Apr":
			hari = 30
		case "Mei":
			hari = 31
		case "Jun":
			hari = 30
		case "Jul":
			hari = 31
		case "Agu":
			hari = 31
		case "Sep":
			hari = 30
		case "Okt":
			hari = 31
		case "Nov":
			hari = 30
		case "Des":
			hari = 31
	default:
		fmt.Println("-")
		return
	}
	
	fmt.Println(hari)
}
```
##### Output
![Screenshot output konversi mil ke km](/praktikum/04-runtunan-sekuensi/tp/soal3/output.png)


#### Deskripsi
Di Soal 3 ini, saya harus menentukan jumlah hari dalam suatu bulan berdasarkan tahun dan nama bulan yang dimasukkan. Bulan yang dimasukkan harus dengan format `3 huruf pertama dari bulan tersebut` dan `Huruf pertamanya harus kapital` Saya menggunakan `switch` untuk menentukan jumlah hari pada setiap bulan (sesuai dengan bulan nya masing-masing). Pada bulan Februari, saya juga membuat supaya program memeriksa jika tahun yang dimasukkan merupakan tahun kabisat, maka Februari berjumlah 29 Hari
```go
case "Feb":
if (tahun%4 == 0 && tahun%100 != 0) || (tahun%400 == 0) {
	hari = 29
} else {
	hari = 28
}
```		
Jika bulan yang dimasukkan tidak sesuai dengan format yang sesuai, yang sudah ada pada `case` yang tersedia di `switch`, maka `-` akan ditampilkan.
```go
default:
	fmt.Println("-")
	return
```

### Soal 4 – Switch Case
`soal4.go`

Membuat satu program Golang yang menerapkan `switch case`

```go
package main
import "fmt"

func main() {
	var (
		makanan []string
		harga []int
		pilihan int
	)

	harga = []int{
		15000,
		5000,
		20000,
		10000,
	}

	makanan = []string{
		"Nasi Goreng",
		"Mi Goreng",
		"Sate",
		"Gado-Gado",
	}

	fmt.Println("============= Makanan =============")
	fmt.Println("Pilih menu makanan")
	for i, m := range makanan {
		fmt.Printf("%d. %s\n", i+1, m)
	}

	fmt.Print("Masukkan pilihan (1-4): ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		fmt.Printf("%s,\nHarga: Rp %d\n", makanan[0], harga[0])
	case 2:
		fmt.Printf("%s,\nHarga: Rp %d\n", makanan[1], harga[1])
	case 3:
		fmt.Printf("%s,\nHarga: Rp %d\n", makanan[2], harga[2])
	case 4:
		fmt.Printf("%s,\nHarga: Rp %d\n", makanan[3], harga[3])
	default:
		fmt.Println("Pilihan tidak valid")
	}
}
```
##### Output
![Screenshot output sisa kue](/praktikum/04-runtunan-sekuensi/tp/soal4/output.png)


#### Deskripsi
Saya membuat program untuk memilih makanan sesuai dengan menu makanan yang tersedia, kemudian menampilkan harganya masing-masing. Program ini dapat menggunakan `switch case`

Daftar harga:
```go
harga = []int{
	15000,
	5000,
	20000,
	10000,
}
```

Daftar makanan:
```go
makanan = []string{
	"Nasi Goreng",
	"Mi Goreng",
	"Sate",
	"Gado-Gado",
}
```

Daftar makanan ditampilkan sehingga pengguna dapat memilih makanan yang diinginkan
```go
fmt.Println("============= Makanan =============")
fmt.Println("Pilih menu makanan")
for i, m := range makanan {
	fmt.Printf("%d. %s\n", i+1, m)
}
```

Memilih makanan
```go
fmt.Print("Masukkan pilihan (1-4): ")
fmt.Scan(&pilihan)
```

Menampilkan makanan dan harganya sesuai dengan makanan yang telah dipilih menggunakan `switch case`
```go
switch pilihan {
case 1:
	fmt.Printf("%s,\nHarga: Rp %d\n", makanan[0], harga[0])
case 2:
	fmt.Printf("%s,\nHarga: Rp %d\n", makanan[1], harga[1])
case 3:
	fmt.Printf("%s,\nHarga: Rp %d\n", makanan[2], harga[2])
case 4:
	fmt.Printf("%s,\nHarga: Rp %d\n", makanan[3], harga[3])
default:
	fmt.Println("Pilihan tidak valid")
}
```

## Kesimpulan
Karena praktikum ini saya dapat lebih memahami bagaimana cara menggunakan `percabangan` atau `if else`, dan `switch case`. Dan juga saya dapat belajar kapan saatnya kita lebih baik menggunakan `if else` dan kapan saatnya kita lebih baik menggunakan `switch case`.
