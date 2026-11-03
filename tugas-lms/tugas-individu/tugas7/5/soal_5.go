package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	naik, turun := true, true
	for i := 1; i < len(s); i++ {
		
		if s[i] <= s[i-1] {
			naik = false
		}

		if s[i] >= s[i-1] {
			turun = false 
		}
	}

	if naik {
		fmt.Println("ascending")
	} else if turun {
		fmt.Println("descending")
	} else {
		fmt.Println("tidak terurut")
	}
}
