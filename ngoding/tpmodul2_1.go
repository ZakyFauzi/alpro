package main
import "fmt"

func reamur(C float64) float64 {
	return 4.0/5.0*C
}

func fahrenheit(C float64) float64 {
	return 9.0/5.0*C+32
}

func kelvin (C float64) float64 {
	return C+273
}

func main() {
	var CAwal, CAkhir, CStep float64
	
	fmt.Scanln(&CAwal, &CAkhir, &CStep)
	
	fmt.Println("Celcius\tReamur\tFahrenheit\tKelvin")
	
	for C := CAwal; C <= CAkhir; C += CStep {
		fmt.Printf("%.2f\t%.2f\t%.2f\t%.2f\n", C, reamur(C), fahrenheit(C), kelvin(C))
	}
}