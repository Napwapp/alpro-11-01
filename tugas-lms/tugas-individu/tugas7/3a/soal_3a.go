package main

import "fmt"

func main() {
	var hari, kenaikan, total, kemarin int

	for {
		var x int
		_, err := fmt.Scan(&x)
		if err != nil || x < 0 || x > 200 {
			break
		}

		hari++
		total += x
		if hari > 1 && x > kemarin {
			kenaikan++
		}
		kemarin = x
	}

	// contoh 3 di slide (masukan "500") berkeluaran "0 0"
	if hari == 0 {
		fmt.Println(0, 0)
		return
	}
	fmt.Printf("%d %.2f\n", kenaikan, float64(total)/float64(hari))
}