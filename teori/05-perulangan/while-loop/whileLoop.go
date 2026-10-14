package main
import "fmt"

func main () {
	var n int

	i := 1
	fmt.Scan(&n)

	for i <= n {
		fmt.Println(i)
		i++
	}
}
