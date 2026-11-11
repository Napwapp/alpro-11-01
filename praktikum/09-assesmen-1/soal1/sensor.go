package main

import "fmt"

func main() {
	var (
		kelembapan, tertinggi, terendah, total, i,lastInput int
		avg                                       float64
	)

	i = 1

	for {
		fmt.Scan(&kelembapan)

		if lastInput == kelembapan {
			break
		}

		if i == 1 {
			tertinggi = kelembapan
			terendah = kelembapan
		}

		if kelembapan > tertinggi {
			tertinggi = kelembapan
		}

		if kelembapan < terendah {
			terendah = kelembapan
		}

		total += kelembapan
		lastInput = kelembapan
		i++
	}

	avg = float64(total) / float64(i-1)
	fmt.Printf("%d %d %.2f", tertinggi, terendah, avg)
}
