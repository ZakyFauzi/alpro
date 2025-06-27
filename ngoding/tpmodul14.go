package main
import "fmt"

const NMAX int = 99
type tabInt [NMAX-1]int

func main() {
	var data tabInt
	var nData int
	
	fmt.Scan(&nData)
	bacaData(&data, nData)
	cetakData(data, nData)
	insertionSort(&data, nData)
	fmt.Println()
	cetakData(data, nData)
}

func insertionSort(A *tabInt, N int) {
    var i, pass, temp int
    
    pass = 1
    for pass < N {
        i = pass
        temp = A[pass]
        for i > 0 && temp < A[i-1] {
            A[i] = A[i-1]
            i = i - 1
        }
        A[i] = temp
        pass = pass + 1
    }
}

func bacaData(A *tabInt, N int) {
    var i int
    
    for i = 0; i < N; i++ {
        fmt.Scan(&A[i])
    }
}

func cetakData(A tabInt, N int) {
    var i int
    
    for i = 0; i < N; i++ {
        fmt.Print(A[i], " ")
    }
}