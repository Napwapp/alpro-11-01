```text
program KapasitasGudang
kamus
   k, v, total, i : integer
algoritma
   input(k)
   total ← 0
   for i ← 1 to 5 do
      input(v)
      total ← total + v
      output(total >= k)
   endfor
endprogram
```