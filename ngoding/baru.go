package main

import "fmt"

func main() {
    var w int
    
    fmt.Scan(&w)
    if (w <= 1 || w > 100 || w%2 != 0) && w != 2{
        fmt.Println("NO")
    } else {
        fmt.Println("YES")
    }
}