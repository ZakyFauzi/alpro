package main
import "fmt"

func lowToUpper(kar rune) rune {
    if kar >= 'a' && kar <= 'z' {
        return kar - 32
    }
    return kar
}

func main() {
    var input rune
    fmt.Scanf("%c", &input)

    output := lowToUpper(input)
    fmt.Printf("%c\n", output)
}