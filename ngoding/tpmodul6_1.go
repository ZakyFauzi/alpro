package main
import "fmt"

type mobil struct {
    merek          string
    tahun_produksi int
    kecepatan      int
}

func main() {
    var m1, m2, m3 mobil
    var rata_rata_kecepatan float64

    fmt.Scan(&m1.merek, &m1.tahun_produksi, &m1.kecepatan)
    fmt.Scan(&m2.merek, &m2.tahun_produksi, &m2.kecepatan)
    fmt.Scan(&m3.merek, &m3.tahun_produksi, &m3.kecepatan)

    rata_rata_kecepatan = float64(m1.kecepatan + m2.kecepatan + m3.kecepatan) / 3.0

    fmt.Printf("Rata-rata kecepatan mobil %s (%d), mobil %s (%d), dan mobil %s (%d): %.2f\n",
	m1.merek, m1.tahun_produksi, m2.merek, m2.tahun_produksi, m3.merek, m3.tahun_produksi,
	rata_rata_kecepatan)
}