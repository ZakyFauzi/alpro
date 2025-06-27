package main
import "fmt"

const NMAX = 10
type tabInt [NMAX]int

func baca(A *tabInt, n *int) {
	var x int
	for *n < NMAX {
		fmt.Scan(&x)
		if x == 0 {
			break
		}
		if x < 0 {
			x = -x
		}
		(*A)[*n] = x
		(*n)++
	}
}

func cetak(A tabInt, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(A[i])
		if i != n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func jumlah(A tabInt, n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += A[i]
	}
	return total
}

func rataRata(A tabInt, n int) float64 {
	if n == 0 {
		return 0
	}
	return float64(jumlah(A, n)) / float64(n)
}

func main() {
	var data tabInt
	var nData int
	baca(&data, &nData)
	cetak(data, nData)
	fmt.Println(jumlah(data, nData), (rataRata(data, nData)))
}