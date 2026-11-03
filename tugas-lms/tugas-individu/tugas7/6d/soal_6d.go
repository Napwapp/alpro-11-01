package main
import "fmt"

func main() {
	var gold, platinum, silver int
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
		} else if poin >= 100 {
			platinum++
		} else {
			silver++
		}
	}

	fmt.Printf("\nGold: %d\n", gold)
	fmt.Printf("Platinum: %d\n", platinum)
	fmt.Printf("Silver: %d\n", silver)
}
