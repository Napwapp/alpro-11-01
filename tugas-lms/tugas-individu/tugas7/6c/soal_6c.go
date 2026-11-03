package main

import "fmt"

func main() {
	var n, gold, platinum, silver, poin int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		poin = 0
		for poin < 50 {
			_, err := fmt.Scan(&poin)
			if err != nil {
				break
			}

			if poin < 50 {
				fmt.Println("Masukkan poin minimal 50")
			}
		}

		if poin > 200 {
			gold++
		} else if poin >= 100 {
			platinum++
		} else {
			silver++
		}
	}

	fmt.Printf("\nGold User: %d\n", gold)
	fmt.Printf("Platinum User: %d\n", platinum)
	fmt.Printf("Silver User: %d\n", silver)
}
