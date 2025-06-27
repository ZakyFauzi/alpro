package main
import "fmt"

const NMAX int = 20
type tabInt [NMAX]int

func main() {
	var A tabInt
	var n int

	baca(&A, &n)
	cetakElemen(A, n)
	cetakInfo(A, n)
}

func baca(A *tabInt, n *int) {
	var x int
	*n = 0

	for *n < NMAX {
		fmt.Scan(&x)
		if x <= 0 {
			break
		}
		A[*n] = x
		*n++
	}
}

func cetakElemen(A tabInt, n int) {
	fmt.Print("Elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf(" %d", A[i])
	}
	fmt.Println()
}

func maksimum(A tabInt, n int) int {
	max := A[0]
	for i := 1; i < n; i++ {
		if A[i] > max {
			max = A[i]
		}
	}
	return max
}

func minimum(A tabInt, n int) int {
	min := A[0]
	for i := 1; i < n; i++ {
		if A[i] < min {
			min = A[i]
		}
	}
	return min
}

func cetakInfo(A tabInt, n int) {
	fmt.Printf("Nilai maksimum: %d\n", maksimum(A, n))
	fmt.Printf("Nilai minimum: %d\n", minimum(A, n))
	fmt.Printf("Banyak elemen: %d\n", n)
}