```text
DEKLARASI
    a, t, m : integer

ALGORITMA:
    INPUT(a, t, m)

    IF (a + t > m) AND (t + m > a) AND (m + a > t) AND (a > 0) AND (t > 0) AND (m > 0) THEN
        PRINT("Segitiga Valid")

        IF (a = t) AND (t = m) THEN
            PRINT("Sama Sisi")
        ELSE IF (a = t) OR (t = m) OR (a = m) THEN
            PRINT("Sama Kaki")
        ELSE
            PRINT("Sembarang")
        ENDIF
    ELSE
        PRINT("Tidak Valid")
    ENDIF
ENDPROGRAM
```