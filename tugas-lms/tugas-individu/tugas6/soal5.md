```text
Algoritma Avatar
Kamus
p, dewasa, kecil, sisa, tidakBerangkat : integer

Algoritma
    input(p)
    tidakBerangkat <- 0
    kecil <- 0

    if p <= 15 then
        dewasa <- p div 5
        if p mod 5 > 0 then dewasa <- dewasa + 1 endif
    else
        dewasa <- 3
        sisa <- p - 15
        if sisa <= 10 then
            kecil <- sisa div 2
            if sisa mod 2 > 0 then kecil <- kecil + 1 endif
        else
            kecil <- 5
            tidakBerangkat <- sisa - 10
        endif
    endif

    output("Dewasa ", dewasa)
    if kecil > 0 then
        output(", kecil ", kecil)
    endif
    
    if tidakBerangkat > 0 then
        output(", dan ", tidakBerangkat, " tak berangkat")
    endif
```