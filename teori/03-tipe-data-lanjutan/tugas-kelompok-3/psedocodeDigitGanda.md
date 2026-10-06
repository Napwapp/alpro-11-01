<!-- Soal 2: Digit Ganda -->
``` 
program soal_Digit_Ganda

kamus
a, loop, hasil, bil1, bil2

algoritma
for loop <- 0 to inf
    input (a)
    if a > 99 then
        output("Masukkan bilangan dua digit saja")
    else then 
endfor

bil1 <- a / 10
bil2 <- a % 10
hasil <- bil1 * 1000 + bil1 * 100 + bil2 * 10 + bil2 
output(hasil)

endprogram
```