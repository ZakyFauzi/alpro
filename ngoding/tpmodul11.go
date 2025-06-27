package main
import "fmt"

const NMAX = 10
type tabInt [NMAX]int

//func SeqSearch(T []int, N int, X int) bool {
//    ketemu := false
//    k := 0
//    for !ketemu && k < N {
//        ketemu = T[k] == X
//        k++
//    }
//    return ketemu
//}

//func SeqSearch(T []int, N int, X int) int {
//    idx := -1
//    k := 0
//    for idx == -1 && k < N {
//        if T[k] == X {
//            idx = k
//        }
//        k++
//    }
//    return idx
//}

func bacaData(A *tabInt, n int) {
    for i := 0; i < n; i++ {
        fmt.Scan(&A[i])
    }
}

func cetakData(A tabInt, n int) {
    fmt.Print("Data Bilangan:")
    for i := 0; i < n; i++ {
        fmt.Print(" ", A[i])
    }
    fmt.Println()
}

func frekuensiBilangan(A tabInt, n, x int) int {
    count := 0
    for i := 0; i < n; i++ {
        if A[i] == x {
            count++
        }
    }
    return count
}

func sequentialSearch(A tabInt, n, x int) bool {
    for i := 0; i < n; i++ {
        if A[i] == x {
            return true
        }
    }
    return false
}

func main() {
    var data tabInt
    var nData, x1 int

    fmt.Scan(&x1)
    fmt.Scan(&nData)    
    if nData > NMAX {
        nData = NMAX
    }
    bacaData(&data, nData)
    cetakData(data, nData)
    fmt.Print("Hasil pencarian: ")
    if sequentialSearch(data, nData, x1) {
        fmt.Printf("Bilangan ditemukan. Terdapat %d bilangan %d.", frekuensiBilangan(data, nData, x1), x1)
    } else {
        fmt.Print("Bilangan tidak ditemukan.")
    }
}