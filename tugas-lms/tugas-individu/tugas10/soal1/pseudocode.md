```text
procedure kemasDonat(in total:integer, in/out p12,p6,p1:integer)
{
    I.S. terdefinisi bilangan bulat yang menyatakan total donat
    F.S. variabel p12, p6 dan p1 berisi banyak paket isi 12, 6 dan 1 dari total donat
}

kamus
    sisa : integer
algoritma
    p12 <- total div 12
    sisa <- total mod 12
    p6 <- sisa div 6
    p1 <- sisa mod 6
endprocedure

program paketDonat
kamus
    total, p12, p6, p1 : integer
algoritma
    input(total)
    kemasDonat(total, p12, p6, p1)
    output(p12, " paket isi 12")
    output(p6, " paket isi 6")
    output(p1, " satuan")
endprogram
```