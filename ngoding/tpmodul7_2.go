package main
import "fmt"

const NMAX = 10
type tabInt [NMAX]int

func bacaData(A *tabInt, n *int) {
	fmt.Scan(n)
	if *n > NMAX {
		*n = NMAX
	}
	for i := 0; i < *n; i++ {
		fmt.Scan(&(*A)[i])
	}
}

func cetakData(A tabInt, n int) {
	if n == 0 {
		fmt.Println("Array kosong")
	} else {
		for i := 0; i < n; i++ {
			fmt.Print(A[i])
			if i != n-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	var data tabInt
	var banyakData int
	bacaData(&data, &banyakData)
	cetakData(data, banyakData)
}