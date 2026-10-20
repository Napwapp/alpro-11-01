package main
import "fmt"

func main () {
	var a string
	fmt.Scan(&a)

	if "a" <= a && a <= "z" ||  "A" <= a && a <= "z" {
		fmt.Println("Alphabet")
	} else {
		fmt.Println("Bukan alphabet")
	}
}
