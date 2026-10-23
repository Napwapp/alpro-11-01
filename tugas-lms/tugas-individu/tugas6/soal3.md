```text
Algoritma Parkir
Kamus
   h1, m1, h2, m2 : integer
   mulaiParkir, selesaiParkir, selisih, hh, mm : integer

Algoritma
   input(h1, m1, h2, m2)

   if h1 <= 5 then h1 <- h1 + 12 endif
   if h2 <= 5 then h2 <- h2 + 12 endif

   mulaiParkir   <- h1 * 60 + m1
   selesaiParkir <- h2 * 60 + m2
   selisih <- selesaiParkir - mulaiParkir

   hh <- selisih div 60
   mm <- selisih mod 60

   output(hh, " jam ", mm, " menit")
```