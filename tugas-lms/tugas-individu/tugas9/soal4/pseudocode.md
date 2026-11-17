```text
function jumlah_kuadrat(n : integer) -> integer
kamus
   i, hasil : integer
algoritma
   hasil <- 0
   i <- 1
   
   for i <= n do
      hasil <- hasil + i * i
      i <- i + 1
   endfor

   return hasil
endfunction

program deret_kuadrat
kamus
   n : integer
algoritma
   input(n)
   output(jumlah_kuadrat(n))
endprogram
```