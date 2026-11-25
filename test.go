package main
import "fmt"

func penjumlahan(a,b int) int {
    return a+b
}

func main() {
    var a,b int
    fmt.Scan(&a, &b)

    fmt.Println(penjumlahan(a, b))
}
