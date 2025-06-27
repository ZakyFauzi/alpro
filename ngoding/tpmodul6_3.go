package main
import "fmt"

type pegawai struct {
    nama       string
    gaji_pokok int
    besar_bonus int
	masa_kerja float64
}

func hitung_bonus(p *pegawai) {
    var pengali float64
    if p.masa_kerja >= 10 {
        pengali = 1.5
    } else if p.masa_kerja >= 5 {
        pengali = 0.75
    } else {
        pengali = 0.5
    }
    p.besar_bonus = int(float64(p.gaji_pokok) * pengali)
}

func main() {
    var p pegawai

    fmt.Scan(&p.nama, &p.gaji_pokok, &p.masa_kerja)
    hitung_bonus(&p)
    fmt.Printf("Pegawai dengan nama %s mendapatkan bonus sebesar Rp %d\n", p.nama, p.besar_bonus)
}