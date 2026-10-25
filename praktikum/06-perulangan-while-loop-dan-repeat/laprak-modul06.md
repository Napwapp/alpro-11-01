# <h1 align="center">Laporan Praktikum Modul 06 - Perulangan While</h1>
<p align="center">Muhamad Nawaf Abduh - 109092630003</p>

## Dasar Teori
### Perulangan While (while-loop)
`while-loop` melakukan pengecekan kondisi diawal, jika kondisi tersebut bernilai `true` maka blok kode di dalamnmya akan dieksekusi, apabila kondisi nya bernilai `false` perulangannya tidak akan dijalankan. 
```text
while (kondisi) do
    // action
endwhile
```
Namun Golang tidak punya `while` atau `do-while`, tapi kita tetap bisa menggunakan konsep yang sama pada kata kunci perulangan yang dimiliki Go yaitu `for`
```go
i := 1
for i <= 5 {
    // action
}
```
Action didalamnya akan terus dijalankan selama kondisi nya bernilai `true`

## Guided
### 1. terkecil.go
```go
package main
import "fmt"

func main() {
    var batas int
    fmt.Scan(&batas)

    n := 1
    for n*n*n <= batas {
        n++
    }

    fmt.Println(n)
}
```

#### Deskripsi
`terkecil.go` Merupakan program untuk mencari nilai bilangan bulat terkecil dari `n` yang dimana ketika `n dipangkatkan dengan 3 > batas` yang dimasukkan

### 2. terbesar.go
```go
package main
import "fmt"

func main() {
    var batas int
    fmt.Scan(&batas)

    n := 1
    for n*n < batas {
        n++
    } 

    fmt.Println(n - 1)
}
```

#### Deskripsi
`terbesar.go` Merupakan program untuk mencari nilai bilangan bulat terbesar dari `n` yang dimana ketika `n kuadrat < dari batas` yang dimasukkan. lalu `n - 1` untuk mengembalikan nilai `n` pada iterasi perulangan terakhir. 

### 3. vokal.go
```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    kalimat, _ := reader.ReadString('\n')

    i := 0
    jumlahVokal := 0

    for i < len(kalimat) {
        ch := strings.ToLower(string(kalimat[i]))
        if ch == "a" || ch == "i" || ch == "u" || ch == "e" || ch == "o" {
            jumlahVokal++
        }
        i++
    }

    fmt.Println(jumlahVokal)
}
```

#### Deskripsi
`vokal.go` Merupakan program untk mencari dan menghitung total ada berapa huruf vokal (a, i, u, e, o) pada suatu kalimat yang dimasukkan. Program menjadikan semua karakter nya ke huruf kecil terlebih dahulu supaya lebih mudah mencari huruf vokal nya.
```go
ch := strings.ToLower(string(kalimat[i]))
if ch == "a" || ch == "i" || ch == "u" || ch == "e" || ch == "o" {
    jumlahVokal++
}
```
jika ditemukan huruf vokal, `jumlahVokal` ditambah 1

## Unguided
### 1. atm.go
```go
package main
import "fmt"

func main() {
	var (
		password string
		pilihan,jumlah int
	)
	const passwordBenar = "1234"

	saldo := 10000000
	i := 0
	loginBerhasil := false

	// Perulangan login: berhenti jika password benar atau i habis
	for !loginBerhasil && i < 3 {
		fmt.Print("Masukkan password ATM: ")
		fmt.Scan(&password)

		if password == passwordBenar {
			loginBerhasil = true
		} else {
			i++
			fmt.Printf("\nPassword salah. Silakan coba lagi.\n")
		}
	}

	// Jika 3 kali salah, ATM terblokir
	if !loginBerhasil {
		fmt.Println("ATM terblokir. Program selesai.")
		return
	}

	fmt.Printf("\nLogin berhasil!\n")

	for pilihan != 3 {
		fmt.Println("\nMenu ATM:")
		fmt.Println("1. Lihat Saldo")
		fmt.Println("2. Tarik Uang")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1: 
			fmt.Printf("\nSaldo Anda: Rp.%d\n", saldo)
		case 2:
			fmt.Print("Masukkan jumlah uang yang ingin ditarik: Rp.")
			fmt.Scan(&jumlah)

			if jumlah <= saldo {
				saldo -= jumlah
				fmt.Println("\nAnda berhasil menarik uang sebesar Rp", jumlah)
				fmt.Println("Saldo tersisa: Rp", saldo)
			} else {
				fmt.Printf("\nSaldo tidak mencukupi\n")
			}
		case 3:
			fmt.Printf("\nKeluar dari sistem. Terima kasih!\n")
		default:
			fmt.Println("\nPilihan tidak valid. Masukkan angka 1 atau 2 atau 3!")
		}
	}
}
```

#### Output
![Screenshot Output Unguided](/praktikum/06-perulangan-while-loop-dan-repeat/unguided/output.png)

#### Output Ketika Passwordnya salah 3x:

![Screenshot Output Unguided](/praktikum/06-perulangan-while-loop-dan-repeat/unguided/output-wrong-password.png)

#### Deskripsi
`atm.go` Merupakan program simulasi sistem ATM. Sebelum masuk ke menu, pengguna memasukkan password terlebih dahulu, dan `password` yang benar adalah `1234`. JIka password yang dimasukkan benar, maka program akan masuk ke menu ATM. Namun jika pengguna salah memasukkan `password` hingga 3x, maka program akan dihentikan dan menampilkan pesan `ATM terblokir. Program selesai.`.
```go
fmt.Println("ATM terblokir. Program selesai.")
```

#### Menu ATM
1. Lihat Saldo
2. Tarik Uang
3. Keluar

#### 1. Menu Lihat Saldo
Menu lihat saldo akan menampilkan saldo yang dimiliki. Pada program ini saldo diinisialisasi dengan `saldo := 10000000`

#### 2. Menu Tarik Uang
`jumlah` penarikan yang dimasukkan akan mengurangi jumlah `saldo` yang ada. Jika penarikannya berhasil program akan mengurangi jumlah `saldo `dengan `jumlah` penarikan yang dimasukkan. dan program akan menampilkan jumlah penarikan dan saldo terbarunya

```go
saldo -= jumlah
fmt.Println("\nAnda berhasil menarik uang sebesar Rp", jumlah)
fmt.Println("Saldo tersisa: Rp", saldo)
```
Jika jumlah penarikan yang di masukkan lebih dari saldo yang ada, maka program akan menampilkan 
`Saldo tidak mencukupi`.

#### 3. Menu Keluar
Program akan dihentikan dan akan menampilkan pesan `Keluar dari sistem. Terima kasih!`

## Kesimpulan
Dari praktikum ini, saya jadi bisa lebih paham untuk menggunakan konsep perulangan `while` di Go menggunakan perulangan `for`. Yaitu kita dapat menggunakan perulangan `for` yang hanya menggunakan kondisi sehingga blok kode didalam perulangannya akan selalu dieksekusi jika kondisi nya bernilai `true`. Cocok untuk studi kasus dimana kita belum tau pasti iterasi perulangannya.

## Referensi
#### 1. Modul 06 Praktikum - Perulangan while