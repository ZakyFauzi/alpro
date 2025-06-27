package main
import "fmt"

func faktorial(n int) int {
    if n == 1 {
        return 1
    }
    return n * faktorial(n-1)
}

func main() {
    var a int
    fmt.Scan(&a)
    fmt.Println(faktorial(a))
}