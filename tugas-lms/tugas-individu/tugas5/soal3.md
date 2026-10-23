```text
program Donat
kamus
   n, m, x, y, donat : integer
algoritma
   input(n, m, x, y)
   donat ← 0
   while n >= x and m >= y do
      n ← n - x
      m ← m - y
      donat ← donat + 1
   endwhile
   output(donat)
endprogram
```