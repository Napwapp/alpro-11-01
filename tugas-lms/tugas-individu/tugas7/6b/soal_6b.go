package main
import "fmt"

func main() {
	poin := 0
	for poin < 50 {
		_, err := fmt.Scan(&poin)
		if err != nil {
			return
		}

		if poin < 50 {
			fmt.Println("Masukkan poin minimal 50")
		}
	}

	if poin > 200 {
		fmt.Println("Gold user")
	} else if poin >= 100 {
		fmt.Println("Platinum user")
	} else {
		fmt.Println("Silver user")
	}
}
