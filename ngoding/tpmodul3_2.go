package main
import "fmt"

func hitungMenang(g, k int, jm *int) {
    if g > k {
        *jm++
    }
}

func hitungDraw(g, k int, jd *int) {
    if g == k {
        *jd++
    }
}

func hitungKalah(g, k int, jk *int) {
    if g < k {
        *jk++
    }
}

func hitungJumGolKegolanSelisih(g, k int, jg, jkg, jsg *int) {
    *jg += g
    *jkg += k
    *jsg = *jg - *jkg
}

func hitungJumPoint(jm, jd int, jp *int) {
    *jp = (jm * 3) + (jd * 1)
}

func main() {
    var n, g, k int
    fmt.Scan(&n)
    jm, jd, jk, jg, jkg, jsg, jp := 0, 0, 0, 0, 0, 0, 0
    for i := 0; i < n; i++ {
        fmt.Scan(&g, &k)
        hitungMenang(g, k, &jm)
        hitungDraw(g, k, &jd)
        hitungKalah(g, k, &jk)
        hitungJumGolKegolanSelisih(g, k, &jg, &jkg, &jsg)
    }
    hitungJumPoint(jm, jd, &jp)
    fmt.Printf("%d %d %d %d %d %d %d %d\n", n, jm, jd, jk, jg, jkg, jsg, jp)
}