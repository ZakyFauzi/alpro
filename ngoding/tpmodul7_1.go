package main
import "fmt"

func main() {
	const NMAX int = 10
	var A [NMAX]string
	var n int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
	for i := 0; i < n; i++ {
		fmt.Println(A[i])
	}
}