## Jawaban
### 1. `25`
### 2. `Nilai akhir result: 25`
### 3. Alur Eksekusi Program
#### Dengan nilai variabel
```go
	x := 10
	y := 5
	z := 15
	result := 0
```

#### Kondisi 1 
```go
if x > 5 { // true
    if y < 10 { // true
        result = x + y // = 15
    } else {
        result = x - y // false
    }
}
```
`x > 5`
`x = 10` maka `x > 5` adalah `true`. Karena kondisi tersebut true maka kondisi selanjutnya (yang berada di dalam `if x > 5 {...}`) dijalankan:
```go
if y < 10 {...}
```
`y < 10`, `y = 5` maka `y < 10` adalah `true`. Karena kondisi tersebut `true` maka
```go
result = x + y
```
dijalankan. Dan, `else{...}` nya tidak dijalankan karena kondisi atau `if` nya  bernilai `true`. Maka sekarang nilai `result` menjadi `15` setelah 
```go 
result = x + y
```
dijalankan.

#### Kondisi 2
```go
if z > 10 && x == 10 { // true
    result += z // = 30
} else {
    result = z - x // false
}
```

- `z > 10` adalah  `true` karena `z = 15`
- `x == 0` adalah `true` karena nilai `x = 10`
- karena kedua kondisi yang dipisahkan menggunakan operator logika `&&` bernilai `true`, maka nilai akhir dari kondisi
`z > 10 && x == 10` adalah `true` sehingga 
```go
    result += z
```
dijalankan. Dan nilai `result` menjadi `30` karena sebelumnya nilai `result = 15` lalu ditambah dengan `z` yang bernilai `15`. `15 + 15 = 30`.

#### Kondisi 3
```go
if x == 10 || y > 10 {
    result += 5 // = 35
} else if y == 5 && z > 10 { 
    result -= 5 // false
} else {
    result *= 2
}
```

-  `x == 10` = true
- `y > 10` = false
- `x == 10 || y > 10` = true

Karena kondisi tersebut `true` maka baris ini dijalankan
```go
result += 5 // = 35
```
Sehingga nilai `result` sekarang menjadi `30` karena ditambah dengan `5`.

`30 + 5 = 35`

Karena kondisi tersebut bernilai `true`, `else if` dan `else` didalamnya tidak akan di eksekusi.

#### Kondisi 4
```go
if !(x < 15 && y < 10) {
    result += 10 // false
} else {
    result -= 10 // 25
}
```
- `!(x < 15 && y < 10)` terdapat tanda `!` yang berarti `NOT` atau `BUKAN` atau `KEBALIKAN` kalau di Logika Matematika disebut `Negasi`. Maka kondisi nya  menjadi `x >= 15 || y >= 10` yang mana hasilnya berniali `false`
- Karena kondisi tersebut bernilai `false` maka baris ini tidak akan dieksekusi
```go
result += 10 // false
```
Karena hal tersebut, `else` akan dieksekusi

```go
else {
    result -= 10 // 25
}
```
- nilai `result` -10
- maka nilai `result` sekarang adalah `25` karena `35 - 10 = 25`
- nilai tersebut merupakan nilai terakhir `result`. Lalu nilai terbarunya ditampilkan menggunakan 
```go
fmt.Println("Nilai akhir result:", result)
```

#### Kesimpulan Soal No. 3
variabel `x`,`y`,dan `z` nilai nya tetap sama hingga akhir. Sedangkan nilai dari variabel `result` terus berubah-rubah sesuai kondisi yang dieksekusi. Dan nilai akhir dari variabel `result` adalah `25`

