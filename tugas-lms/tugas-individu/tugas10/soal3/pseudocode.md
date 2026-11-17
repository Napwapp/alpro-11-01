```text
procedure mengurutkanMenurun(in/out a,b:integer)
{
    I.S. terdefinisi bilangan bulat a dan b
    F.S. nilai a > b, proses pertukaran nilai pada variabel a dan b mungkin dilakukan
}

kamus
    temp : integer
algoritma
    if a < b then
        temp <- a
        a <- b
        b <- temp
    endif
endprocedure

program terurutMenurun
kamus
    x, y : integer
algoritma
    input(x, y)
    while x ≠ y do
        mengurutkanMenurun(x, y)
        output(x, " ", y)
        input(x, y)
    endwhile
endprogram
```