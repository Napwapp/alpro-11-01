```text
procedure konversiBerat(in kg:real, in/out gram,ons,pon:real)
{
    I.S. terdefinisi bilangan riil kg yang menyatakan berat dalam kilogram
    F.S. variabel gram, ons dan pon berisi hasil konversi dari kg
}

algoritma
    gram <- kg * 1000
    ons <- gram / 100
    pon <- gram / 500
endprocedure

program konversiBerat
kamus
    kg, gram, ons, pon : real
algoritma
    input(kg)
    konversiBerat(kg, gram, ons, pon)
    output(gram, "g ", ons, " ons ", pon, " pon")
endprogram
```