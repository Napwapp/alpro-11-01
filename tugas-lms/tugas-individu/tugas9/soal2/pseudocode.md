```text
function letak_titik(x, y : integer) -> string
kamus
   hasil : string
algoritma
   if x == 0 and y == 0 then
      hasil <- "Titik Asal"
   else if y == 0 then
      hasil <- "Sumbu X"
   else if x == 0 then
      hasil <- "Sumbu Y"
   else if x > 0 and y > 0 then
      hasil <- "Kuadran I"
   else if x < 0 and y > 0 then
      hasil <- "Kuadran II"
   else if x < 0 and y < 0 then
      hasil <- "Kuadran III"
   else
      hasil <- "Kuadran IV"
   endif
   return hasil
endfunction

program kuadran
kamus
   x, y : integer
algoritma
   input(x, y)
   output(letak_titik(x, y))
endprogram
```