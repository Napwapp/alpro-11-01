```text
Algoritma AkhirTahun
Kamus
   total, bayar : integer
   bersedia, kartu, diskon, cashback : boolean

Algoritma
   input(total, bersedia)

   diskon   <- total >= 100000
   kartu    <- bersedia and diskon
   cashback <- total >= 200000 and kartu

   
   if diskon then
      total <- total - 0.1 * total
   endif
   if cashback then
      total <- total - 75000
   endif

   output("Kartu? ", kartu)
   output("Diskon? ", diskon)
   output("Cashback? ", cashback)
   output("Rp. ", total)
```