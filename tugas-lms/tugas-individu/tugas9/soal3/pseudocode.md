```text
function tarif(jam : integer) -> integer
kamus
   total : integer
algoritma
   if jam <= 2 then
      total <- 5000
   else
      total <- 5000 + (jam - 2) * 3000
   endif
   return total
endfunction

program tarif_parkir
kamus
   jam : integer
algoritma
   input(jam)
   output(tarif(jam))
endprogram
```