// function binSearch(tab tabInt, n int)bool{
// 	var left, right, mid int
	
// 	left = 0
// 	right = n-1
// 	mid = (left + right) / 2
// 	for left <= right && tab[mid] != x {
// 		if x < tab[mid] {
// 			right = mid - 1
// 		}
// 		else {
// 			left = mid + 1
// 		}
// 		mid = (left + right) / 2
// 	}
// 	return mid > -1 && tab[mid] == x
// }


// function binSearch(T tab, n, x int)int{
// 	var left, right, mid, idx int
	
// 	left = 0
// 	right = n-1
// 	idx = -1
// 	for left <= right && idx == -1 {
// 		mid = (left + right) div 2
// 		if x < tab[mid] {
// 			right = mid - 1
// 		} else if x > tab[mid] {
// 			left = mid + 1
// 		} else {
// 			idx = mid
// 		}
// 	}
// 	return idx
// }


package main
import "fmt"

const NMAX = 10
type tabInt [NMAX]int

func main() {
    var data tabInt
    var idx, nData, x1 int
    
    fmt.Scan(&nData)
    if nData > NMAX {
        nData = NMAX
    }
    bacaData(&data, nData)
    fmt.Scan(&x1)
    cetakData(data, nData)
    idx = binarySearch(data, nData, x1)
    if idx != -1 {
        fmt.Print("Data ditemukan pada indeks ke-", idx)
    } else {
        fmt.Println("Data tidak ditemukan")
    }
}

func bacaData(A *tabInt, n int) {
    var i int
	
	for i = 0; i < n && i < NMAX; i++ {
        fmt.Scan(&A[i])
    }
}

func cetakData(A tabInt, n int) {
    var i int
	
	fmt.Println("Data dalam array:")
    for i = 0; i < n; i++ {
        fmt.Print(A[i])
        if i < n-1 {
            fmt.Print(" ")
        }
    }
    fmt.Println()
}

func binarySearch(A tabInt, n, x int) int {
    var left, right, mid, idx int
	left = 0
    right = n - 1
    idx = -1
    
    for left <= right && idx == -1 {
        mid = (left + right) / 2
        if x < A[mid] {
            right = mid - 1
        } else if x > A[mid] {
            left = mid + 1
        } else {
            idx = mid
        }
    }    
    return idx
}