package main 
import "fmt"

func main() {
	var (
		n, lebarDaun, min int
	) 

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&lebarDaun)

		if i == 0 {
			min = lebarDaun
		} else if lebarDaun < min {
			min = lebarDaun
		}
	}
	fmt.Println(min)
}