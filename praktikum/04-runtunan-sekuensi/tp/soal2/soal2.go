// Soal 2 – Tracing: Evaluasi Pernyataan Kondisi
package main

import "fmt"

func main() {
	x := 10
	y := 5
	z := 15
	result := 0

	if x > 5 { // true
		if y < 10 { // true
			result = x + y // = 15
		} else {
			result = x - y // false
		}
	}

	if z > 10 && x == 10 { // true
		result += z // = 30
	} else { // false
		result = z - x 
	}

	if x == 10 || y > 10 {
		result += 5 // = 35
	} else if y == 5 && z > 10 { 
		result -= 5 // false
	} else {
		result *= 2
	}

	if !(x < 15 && y < 10) {
		result += 10 // false
	} else {
		result -= 10 // 25
	}

	fmt.Println("Nilai akhir result:", result)
}
