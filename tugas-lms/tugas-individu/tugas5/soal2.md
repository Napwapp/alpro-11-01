```text
program KuadratDigit
kamus
   x, digit, kuadrat, total : integer
   hasil : string
algoritma
   input(x)
   total ← 0
   hasil ← ""
   while x > 0 do
      digit ← x mod 10
      kuadrat ← digit * digit
      if hasil != "" then
         hasil ← hasil + " "
      endif
      hasil ← hasil + toString(kuadrat)
      total ← total + kuadrat
      x ← x div 10
   endwhile
   output(hasil)
   output(total)
endprogram
```