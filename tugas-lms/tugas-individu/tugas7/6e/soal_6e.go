package main

import "fmt"

func main() {
	var gold, platinum, silver, sumGold, sumPlatinum, sumSilver int
	total := 0

	for total < 500 {
		poin := 0
		for poin < 50 {
			_, err := fmt.Scan(&poin)
			if err != nil {
				break
			}

			if poin < 50 {
				fmt.Println("Masukkan poin minimal 50")
			}
		}

		total += poin
		if poin > 200 {
			gold++
			sumGold += poin
		} else if poin >= 100 {
			platinum++
			sumPlatinum += poin
		} else {
			silver++
			sumSilver += poin
		}
	}

	avgGold, avgPlatinum, avgSilver := 0.0, 0.0, 0.0
	if gold > 0 {
		avgGold = float64(sumGold) / float64(gold)
	}
	if platinum > 0 {
		avgPlatinum = float64(sumPlatinum) / float64(platinum)
	}
	if silver > 0 {
		avgSilver = float64(sumSilver) / float64(silver)
	}

	fmt.Printf("Gold: %.2f\n", avgGold)
	fmt.Printf("Platinum: %.2f\n", avgPlatinum)
	fmt.Printf("Silver: %.2f\n", avgSilver)
}
