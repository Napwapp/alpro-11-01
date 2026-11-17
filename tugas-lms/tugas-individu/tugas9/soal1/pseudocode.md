```text
function konversi(km,m,cm : integer) -> integer 
kamus
    hasil : integer
algoritma
    hasil <- km * 100000 + m * 100 + cm
   return hasil
endfunction

program konversi_jarak
kamus
    km,m,cm : integer
algoritma
    input(km,m,cm)
    output("Hasil konversi = ", konversi(km, m, cm), " cm")
endprogram
```