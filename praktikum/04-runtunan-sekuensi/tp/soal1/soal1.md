### Kondisi Yang Dipilih
```go
intOther > 0 && intNum > 0 || sngNum > 0
```
### Kode program
```go
package main
import "fmt"

func main () {
	intNum := 5
	intOther := 10
	var sngNum float64 = -3
	
	if intOther > 0 && intNum > 0 || sngNum > 0 {
		fmt.Println("Beep")
	} 
}
```

### Output
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/tp/soal1/output.png)

### Jawaban

Dengan nilai `intNum = 5`, `intOther = 10`, dan `sngNum = -3`:

1. `intNum > 5` = **false**
2. `intNum >= 5 && intOther < 11` = **true**
3. `sngNum != -1 || intOther < 0` = **true**
4. `!(intNum > 3) || intNum <= 5` = **true**
5. `!(intOther >= intNum)` = **false**
6. `0 - sngNum > 0` = **true**
7. `4 / 2 == intOther / intNum` = **true**
8. `intOther % 2 == 0` = **true**
9. `intOther + 2 * intNum != 30 || !(sngNum > 0)` = **true**
10. `intOther > 0 && intNum > 0 || sngNum > 0` = **true**
11. `sngNum > 0 || (intNum >= 0 && -1 * intOther == -10)` = **true**
12. `intNum == 5` = **true**
13. `intNum > 0 || (sngNum <= 0 && intOther == 13)` = **true**
14. `!(!(!(!(intNum > 0))))` = **true**