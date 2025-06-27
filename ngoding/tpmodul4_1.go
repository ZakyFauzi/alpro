package main
import "fmt"

func fibonacci(n int) int {
    if n == 1 || n == 2 {
        return 1
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    var a int
    fmt.Scan(&a)
    fmt.Println(fibonacci(a))
}