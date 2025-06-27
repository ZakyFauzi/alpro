package main
import "fmt"

func persamaan(a, b int) int {
    if b == 0 {
        return a
    }
    return persamaan(b, a%b)
}

func main() {
    var x, y int
    fmt.Scan(&x, &y)
    fmt.Println(persamaan(x, y))
}