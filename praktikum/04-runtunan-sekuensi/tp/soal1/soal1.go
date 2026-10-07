// Soal 1 – Evaluasi Ekspresi Kontrol dalam Go
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