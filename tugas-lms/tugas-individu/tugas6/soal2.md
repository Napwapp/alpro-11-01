```text
Algoritma GolLiga
Kamus
   k1, k2, k3, k4, terbanyak, tersedikit : integer

Algoritma
    input(k1, k2, k3, k4)
    terbanyak  <- k1
    tersedikit <- k1

    if k2 > terbanyak then
        terbanyak <- k2
    endif
   
    if k3 > terbanyak then 
        terbanyak <- k3
    endif
    
    if k4 > terbanyak then 
        terbanyak <- k4
    endif   
    
    if k2 < tersedikit then 
        tersedikit <- k2
    endif
    
    if k3 < tersedikit then 
        tersedikit <- k3
    endif
    
    if k4 < tersedikit then 
        tersedikit <- k4
    endif

   output(terbanyak, tersedikit)
```

