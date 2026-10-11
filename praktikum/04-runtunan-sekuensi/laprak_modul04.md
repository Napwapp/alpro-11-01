# <h1 align="center">Laporan Praktikum Modul 04 - Runtunan atau Sekuensi</h1>
<p align="center">Muhamad Nawaf Abduh - 109092630003</p>

## Dasar Teori

###  Pengertian Runtunan atau Sekuensi
Runtunan atau Sekuensi merupakan kata lain dari `berurutan`.

Dalam pemrograman, instruksi yang kita buat menggunakan sebuah bahasa pemrograman (apapun itu), akan dieksekusi oleh komputer secara berurutan dari instruksi pertama hingga terakhir atau dari `atas` ke `bawah` dan dari `kiri` ke `kanan`. Mengetahui bagaimana runtunan yang baik saat kita membuat suatu program itu sangat penting. Karena, apabila kita membuat sebuah instruksi  atau baris kode yang urutannya kurang tepat, itu bisa memengaruhi hasil akhir dari program yang dibuat.

#### Contoh: 
```go
package main
import "fmt"
func main() {
 nama := "Ahmad Fauzan"
 nim := "1301230012"
 jurusan := "S1 Informatika"
 
 fmt.Println("=== BIODATA MAHASISWA ===")
 fmt.Println("Nama :", nama)
 fmt.Println("NIM :", nim)
 fmt.Println("Jurusan :", jurusan)
}
```
bila program ini dijalankan, maka outputnya:
```text 
=== BIODATA MAHASISWA ===
Nama : Ahmad Fauzan
NIM : 1301230012
Jurusan : S1 Informatika
```
lalu, apa yang akan terjadi jika kita salah megurutkan penulisannya?? Misalnya kita menuliskan program nya seperti ini (**misal saja**):
```go
package main
import "fmt"
func main() { 
	fmt.Println("=== BIODATA MAHASISWA ===")
	fmt.Println("Nama :", nama)
	fmt.Println("NIM :", nim)
	fmt.Println("Jurusan :", jurusan)

	nama := "Ahmad Fauzan"
	nim := "1301230012"
	jurusan := "S1 Informatika"
}
```
Jika program tersebut dijalankan, **program tidak akan bisa dieksekusi** karena variabel `nama`, `nim`, dan `jurusan` yang digunakan pada perintah `fmt.Println` belum dideklarasikan (Ingat! program yang kita buat, dieksekusi secara berurutan dari atas ke bawah dan kiri ke kanan). Karena deklarasi variabel `nama`, `nim` dan `jurusan` nya berada di bawah setelah `fmt.Println()` yang memanggil ketiga `variabel` tersebut, komputer tidak bisa mengenali variabel `nama`, `nim`, dan `jurusan` karena ketiga `variabel` tersebut baru di deklarasikan setelah ketiga `fmt.Println()` itu dijalankan, sehingga akan menghasilkan **error**

Compiler akan menampilkan pesan kesalahan seperti:
- `undefined: nama`, 
- `undefined: nim`, 
- `undefined: jurusan`.

Urutan penulisan instruksi pada sebuah program, sangat berpengaruh pada hasil akhir atau output dari program yang dibuat.

## Guided
### 1. grade.go

```go
package main
import "bufio"
import "fmt"
import "os"

func main() {
	var (
		nilai float64
		nama,grade string
	)
	// bufio scanner untuk input
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("Masukkan nama: ")
	scanner.Scan()
	nama = scanner.Text()

	fmt.Println("Masukkan nilai: ")
	fmt.Scan(&nilai)

	if nilai >= 90 && nilai <= 100 {
		grade = "A"
	} else if nilai >= 80 && nilai < 90 {
		grade = "B"
	} else if nilai >= 70 && nilai < 80 {
		grade = "C"
	} else if nilai >= 60 && nilai < 70 {
		grade = "D"
	} else {
		grade = "E"
	}

fmt.Println(nama)
fmt.Println(grade)}
```
#### Deskripsi
`grade.go` merupakan program untuk memberikan klasifikasi pada nilai siswa yang diinputkan berdasarkan sistem penilaian yang ada.

Pertama, kita menginputkan nama dan nilai siswa terlebih dahulu.
```go
// bufio scanner untuk input
scanner := bufio.NewScanner(os.Stdin)

fmt.Println("Masukkan nama: ")
scanner.Scan()
nama = scanner.Text()

fmt.Println("Masukkan nilai: ")
fmt.Scan(&nilai)
```
menggunakan `bufio` supaya kita bisa menginputkan `string` lebih dari 1 kata.

Kemudian, program ini menggunakan pengkondisian bercabang supaya dapat mengklasifikan nilai siswa yang diinputkan sesuai dengan rentang nilai dan sistem penilaian nya
```go
if nilai >= 90 && nilai <= 100 {
	grade = "A"
} else if nilai >= 80 && nilai < 90 {
	grade = "B"
} else if nilai >= 70 && nilai < 80 {
	grade = "C"
} else if nilai >= 60 && nilai < 70 {
	grade = "D"
} else {
	grade = "E"
}
```

Sistem Penilaian :
1. A = Apabila nilai siswa `90` sampai `100`
2. B = Apabila nilai siswa `80` sampai sampai kurang dari `90`
3. C = Apabila nilai siswa `70` sampai kurang dari `80`
4. D = Apabila nilai siswa `60` sampai kurang dari `70`
5. E = Apabila nilai siswa kurang dari `60`

Kemudian, outputnya ditampilkan:
```go
fmt.Println(nama)
fmt.Println(grade)
```

#### Contoh Input, Output
#### Input
```text
Muhamad Nawaf Abduh
95
```
#### Output
```text
Muhamad Nawaf Abduh
A
```

### 2. penilaian.go
```go
package main

import "bufio"
import "fmt"
import "os"

func main() {
var (
	nama, grade    string
	nilai, pilihan int
)

fmt.Println("============= Menu =============")
fmt.Println("1. Sistem penilaian")
fmt.Println("0. Keluar")
fmt.Print("Pilih opsi: ")
fmt.Scanln(&pilihan)

if pilihan == 1 {
	fmt.Print("Masukkan Nama:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nama = scanner.Text()

	fmt.Print("Masukkan Nilai:")
	fmt.Scan(&nilai)

	if nilai >= 90 && nilai <= 100 {
		grade = "A"
	} else if nilai >= 80 && nilai < 90 {
		grade = "B"
	} else if nilai >= 70 && nilai < 80 {
		grade = "C"
	} else if nilai >= 60 && nilai < 70 {
		grade = "D"
	} else {
		grade = "E"
	}

	fmt.Println(nama)
	fmt.Println(grade)

} else if pilihan == 0 {
	return
} else {
	fmt.Println("Pilihan tidak valid")
}
}
```

#### Deskripsi
`penilaian.go` merupakan program sistem penilain berdasarkan menu pilihan yaitu `1` atau `0`.
```go
fmt.Println("============= Menu =============")
fmt.Println("1. Sistem penilaian")
fmt.Println("0. Keluar")
fmt.Print("Pilih opsi: ")
fmt.Scanln(&pilihan)
```

- Jika memasukkan angka 1: program akan melanjutkan ke sistem penilaian
- Jika memasukkan angka 0: program akan diberhentikan
- Jika memasukkan angka lain: Akan menampilkan `"Pilihan tidak valid"`
```go
else {
	fmt.Println("Pilihan tidak valid")
}
```

#### Memasukkan angka 1 atau Melanjutkan ke Sistem Penilaian
```go 
if pilihan == 1 {
	fmt.Print("Masukkan Nama:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nama = scanner.Text()

	fmt.Print("Masukkan Nilai:")
	fmt.Scan(&nilai)

	if nilai >= 90 && nilai <= 100 {
		grade = "A"
	} else if nilai >= 80 && nilai < 90 {
		grade = "B"
	} else if nilai >= 70 && nilai < 80 {
		grade = "C"
	} else if nilai >= 60 && nilai < 70 {
		grade = "D"
	} else {
		grade = "E"
	}

	fmt.Println(nama)
	fmt.Println(grade)
}
```
Jika kamu memasukkan angka 1, program akan melanjutkan ke sistem penilaian dengan menginputkan `nama` dan `nilai` siswa lalu pengkondisian bercabang pada program akan mengklasifikasi kan `grade` apa yang akan didapat oleh siswa tersebut sesuai rentang nilai nya (Rentang nilai nya sama saja dengan yang di soal nomor 1).

#### Contoh Input Output
#### Input
```text
Devand
57
```

#### Output
```text
Devand
E
```

#### Memasukkan angka 0
Jika kamu memasukkan angka 0, maka program akan langsung dihetikan

#### Jika memeasukkan angka selain 1 dan 0
Akan menampilkan `"Pilihan tidak valid"`
```go
else {
	fmt.Println("Pilihan tidak valid")
}
```

### 3. klasifikasi.go
```go
package main
import "fmt"

func main() {
	var (
		usia,gajiTahunan int
	)

	fmt.Print("Masukkan usia: ")
	fmt.Scan(&usia)

	fmt.Print("Masukkan gaji tahunan (juta): ")
	fmt.Scan(&gajiTahunan)

	switch {
	case usia < 18 && gajiTahunan == 0 :
		fmt.Println("Masih Sekolah")
	case (usia >= 18 && usia <= 25) && gajiTahunan >= 50 :
		fmt.Println("Muda Sukses")
	case (usia >= 18 && usia <= 25) && gajiTahunan < 50 :
		fmt.Println("Masih belajar hidup")
	case (usia >= 26 && usia <= 40) && gajiTahunan >= 100 :
		fmt.Println("Pekerja Mapan")
	case (usia >= 26 && usia <= 40) && gajiTahunan < 100 :
		fmt.Println("Perlu perbaikan karir")
	case usia > 40 && gajiTahunan >= 150 :
		fmt.Println("Profesional Berpengalaman")
	case usia > 40 && gajiTahunan < 150 :
		fmt.Println("Perlu Evaluasi Finansial")
	}
}
```

#### Deskripsi
Berbeda dari 2 tugas guided sebelumnya yang menggunakan pengkondisian bercabang, di `klasifikasi.go` menggunakan `switch cases`.

`klasifikasi.go` Merupakan program untuk mengklasifikan status seseorang berdasarkan `usia`, dan total `gaji tahunan` nya.
```go
switch {
case usia < 18 && gajiTahunan == 0 :
	fmt.Println("Masih Sekolah")
case (usia >= 18 && usia <= 25) && gajiTahunan >= 50 :
	fmt.Println("Muda Sukses")
case (usia >= 18 && usia <= 25) && gajiTahunan < 50 :
	fmt.Println("Masih belajar hidup")
case (usia >= 26 && usia <= 40) && gajiTahunan >= 100 :
	fmt.Println("Pekerja Mapan")
case (usia >= 26 && usia <= 40) && gajiTahunan < 100 :
	fmt.Println("Perlu perbaikan karir")
case usia > 40 && gajiTahunan >= 150 :
	fmt.Println("Profesional Berpengalaman")
case usia > 40 && gajiTahunan < 150 :
	fmt.Println("Perlu Evaluasi Finansial")
}
```

#### Contoh Input, Output.
#### Input
```text
usia: 21
gaji tahunan: 52
```

#### Output
```text
Muda Sukses
```

## Unguided

### 1. grade.go dengan menggunakan switch case
```go
package main
import "bufio"
import "fmt"
import "os"

func main() {
	var (
		nilai float64
		nama,grade string
	)
	// bufio scanner untuk input
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("Masukkan nama: ")
	scanner.Scan()
	nama = scanner.Text()

	fmt.Println("Masukkan nilai: ")
	fmt.Scan(&nilai)

	switch {
	case nilai >= 90 && nilai <= 100:
		grade = "A"
	case nilai >= 80 && nilai < 90:
		grade = "B"
	case nilai >= 70 && nilai < 80:
		grade = "C"
	case nilai >= 60 && nilai < 70:
		grade = "D"
	default:
		grade = "E"
	}

	fmt.Printf("Nama: %s\nGrade: %s", nama, grade)
}
```

#### Output
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/unguided/grade/output.png)

#### Deskripsi
Program ini sama saja dengan program `grade.go` pada tugas guided sebelumnya. Namun bedanya, di sini menggunakan switch cases untuk mengklasisfikan grade nya berdasarkan rentang nilai nya masing-masing. Lalu ditampilkan menggunakan
```go
	fmt.Printf("Nama: %s\nGrade: %s", nama, grade)
```
dengan menggunakan sedikit formatting.
### 2. pajak.go

```go
package main

import "fmt"

func main() {
	var (
		penghasilan, pajak float64
	)

	// Percents
	p5 := 5.0 / 100.0
	p10 := 10.0 / 100.0
	p15 := 15.0 / 100.0
	p20 := 20.0 / 100.0

	fmt.Print("Masukkan penghasilan (juta): ")
	fmt.Scan(&penghasilan)

	switch {
	case penghasilan <= 50:
		pajak = penghasilan * p5
		fmt.Printf("Total pajak: %.1f juta\n", pajak)

	case penghasilan > 50 && penghasilan <= 100:
		pajak = (50 * p5) + ((penghasilan - 50) * p10)
		fmt.Printf("Total pajak: %.1f juta\n", pajak)

	case penghasilan > 100 && penghasilan <= 200:
		pajak = (50 * p5) + (50 * p10) + ((penghasilan - 100) * p15)
		fmt.Printf("Total pajak: %.1f juta\n", pajak)

	case penghasilan > 200:
		pajak = (50 * p5) + (50 * p10) + (100 * p15) + ((penghasilan - 200) * p20)
		fmt.Printf("Total pajak: %.1f juta\n", pajak)
	}
}
```

#### Output
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/unguided/pajak/output.png)


#### Deskripsi
`pajak.go` merupakan program untuk menghitung pajak yang harus dibayarkan seseorang berdasarkan penghasilan dalam juta rupiah. Perhitungannya menggunakan beberapa lapisan: persentase pajak yang lebih tinggi hanya dikenakan pada bagian penghasilan yang masuk ke lapisan tersebut, bukan pada seluruh penghasilan.

1. Jika penghasilan maksimalnya `50 juta`, dia harus membayar pajak sebesar `5%` dari penghasilannya.
2. Jika penghasilannya lebih dari `50 juta` sampai `100 juta`, pajak `5%` dikenakan pada `50 juta` pertama. Sisa penghasilannya setelah dikurang `50 juta pertama` tadi dikenakan pajak sebesar `10%`.
3. Jika penghasilannya lebih dari `100 juta` sampai `200 juta`, dikenakan pajak `5%` untuk `50 juta pertama`, `10%` untuk `50 juta berikutnya`, lalu sisa penghasilannya yang belum dikenakan pajak, dikenakan pajak sebesar `15%`.
4. Jika penghasilannya lebih dari `200 juta`, pajak dihitung dari `5%` untuk `50 juta pertama`, `10%` untuk `50 juta berikutnya`, `15%` untuk `100 juta berikutnya`, lalu sisa penghasilannya yang belum dihitung pajak, dikenakan pajak sebesar `20%`.

Lalu, program menampilkan total pajak dengan format satu angka di belakang koma.

```go
fmt.Printf("Total pajak: %.1f juta\n", pajak)
```

## Kesimpulan
Dalam pemrograman, instruksi yang kita buat menggunakan bahasa pemrograman apapun itu, akan dieksekusi secara berurutan dari instruksi pertama hingga akhir, dari atas ke bawah, dan dari kiri ke kanan. Kita harus bisa menentukan urutan penulisan instruksi yang tepat supaya hasil akhir dari program yang kita buat bisa sesuai dengan apa yang kita inginkan

## Referensi
#### 1. Modul 04 Praktikum - Runtunan atau Sekuensi