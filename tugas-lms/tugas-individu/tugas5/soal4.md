```text
program BilanganCantik
kamus
   x, kanan, kiri : integer
   cantik : boolean
algoritma
   input(x)
   cantik ← true
   kanan ← x mod 10
   x ← x div 10
   while x > 0 and cantik do
      kiri ← x mod 10
      if kiri > kanan then
         cantik ← false
      endif
      kanan ← kiri
      x ← x div 10
   endwhile
   output(cantik)
endprogram
```