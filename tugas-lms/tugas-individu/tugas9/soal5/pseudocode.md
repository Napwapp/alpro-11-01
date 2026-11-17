```text
function faktorial(n : integer) -> integer
kamus
   i, hasil : integer
algoritma
   hasil <- 1
   i <- 1
   for i <= n do
      hasil <- hasil * i
      i <- i + 1
   endfor
   return hasil
endfunction

function kombinasi(n, r : integer) -> integer
kamus
   hasil : integer
algoritma
   hasil <- faktorial(n) / (faktorial(r) * faktorial(n - r))
   return hasil
endfunction

program faktorial_kombinasi
kamus
   x, y, c : integer
algoritma
   input(x, y)
   if x >= y then
      c <- kombinasi(x, y)
   else
      c <- kombinasi(y, x)
   endif
   output(faktorial(x), " ", faktorial(y), " ", c)
endprogram
```