package main
import "fmt"

const NMAX = 5
type tabInt struct {
	info [NMAX]int
	n    int
}

func bacaData(A *tabInt) {
	var input int
	fmt.Scan(&input)
	if A.n < NMAX {
		A.info[A.n] = input
		A.n++
	}
}

func cetakData(A tabInt) {
	for i := 0; i < A.n; i++ {
		fmt.Print(A.info[i])
		if i != A.n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func main() {
	var data tabInt
	for i := 0; i < 6; i++ {
		bacaData(&data)
	}
	cetakData(data)
}