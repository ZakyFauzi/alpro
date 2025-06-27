package main
import "fmt"

func adaBilanganM(n, m int) string {
    for n > 0 {
        if n%10 == m {
            return "YA"
        }
        n /= 10
    }
    return "TIDAK"
}

func main() {
    var N, M int

    fmt.Scanln(&N, &M)
	
    ada := adaBilanganM(N, M)
    fmt.Println(ada)
}