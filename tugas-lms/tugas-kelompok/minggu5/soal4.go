package main

import "fmt"

func main (){
	var a, t, m int // a untuk alas, t sisi 1, m untuk sisi 2
	fmt.Println("====== Segitiga ======")
	fmt.Print("Alas: ")
	fmt.Scan(&a)
	fmt.Print("Sisi_Satu: ")
	fmt.Scan(&t)
	fmt.Print("Sisi_Dua: ")
	fmt.Scan(&m)

	//percabangan
	if (a + t > m) && (t + m > a) && (m + a > t) && ((a > 0) && (t > 0) && (m > 0)){
		fmt.Println("Segitiga Valid")
		if (a == t) && (t != m) || (t == m) && (m != a) || (m == a) && (a != t){
			fmt.Print("Segitiga Sama Kaki")
		} else if (a*a + t*t == m*m) || (a*a + m*m == t*t) {
			fmt.Print("Segitiga Siku-Siku")
		} else if (a == t) && (t == m) {
			fmt.Print("Segitiga Sama Sisi")
		} else {
			fmt.Print("Segitiga Sembarangan")
		}
	} else {
		fmt.Printf("Segitiga Tidak Valid")
	}
}