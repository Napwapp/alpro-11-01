```text
Algoritma Profit
Kamus
   bulanLalu, bulanIni : real

Algoritma
   input(bulanLalu, bulanIni)
   
   if bulanIni > bulanLalu then
      output("Peningkatan sebesar ", bulanIni - bulanLalu)
   else if bulanIni < bulanLalu then
      output("Penurunan sebesar ", bulanLalu - bulanIni)
   else
      output("Tetap")
   endif
```