package main
import "fmt"

type pemain [1000]string
type tabInt [1000]int

func main() {
	var n, i int
	var p pemain
	var set1, set2, set3, total tabInt

	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&p[i], &set1[i], &set2[i], &set3[i])
	}

	for i = 0; i < n; i++ {
		total[i] = set1[i] + set2[i] + set3[i]
		rata := float64(total[i]) / 3

		if rata < 40 {
			fmt.Println(p[i], "perlu ikut latihan ekstra selama 1 minggu")
		} else if rata >= 40 && rata < 50 {
			fmt.Println(p[i], "perlu ikut latihan ekstra selama 3 hari")
		} else {
			fmt.Println(p[i], "tidak perlu mengikuti latihan tambahan")
		}
	}
}
