package main
import "fmt"

const pi = 3.14

func hitungLuasKelilingLingkaran(r float64) (float64, float64) {
	luas := pi * r * r
	keliling := 2 * pi * r
	return luas, keliling
}

func hitungLuasKelilingPersegi(s float64) (float64, float64) {
	luas := s * s
	keliling := 4 * s
	return luas, keliling
}

func hitungTotal(lL, lP, kL, kP float64) (float64, float64) {
	return lL + lP, kL + kP
}

func main() {
	var r, s float64
	var firstInput = true
	for {
		fmt.Scan(&r, &s)
		if r == 0 && s == 0 {
			break
		}
		ll, kl := hitungLuasKelilingLingkaran(r)
		lp, kp := hitungLuasKelilingPersegi(s)
		tl, tp := hitungTotal(ll, lp, kl, kp)
		if firstInput {
			fmt.Println("R       S      LL      LP      KL      KP      TL      TP")
			firstInput = false
		}
		fmt.Printf("%6.2f %6.2f %6.2f %6.2f %6.2f %6.2f %6.2f %6.2f\n", r, s, ll, lp, kl, kp, tl, tp)
	}
}
