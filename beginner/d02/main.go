package main

import "fmt"

const (
	ktm = 0.621
	mtk = 1 / ktm
	ktl  = 2.2
	ltk  = 1 / ktl
)

func kmToMiles(km float64) float64 {
	return km * ktm
}
func milesToKm(miles float64) float64 {
	return miles * mtk
}
func kgsToLbs(kgs float64) float64 {
	return kgs * ktl
}
func lbsToKgs(lbs float64) float64 {
	return lbs * ltk
}
func cToF(c float64) float64 {
	return (c * 1.8) + 32
}
func fToC(f float64) float64 {
	return (f - 32) * 0.55
}
func main() {
	fmt.Println("Km to Miles",kmToMiles(10))
	fmt.Println("Miles to Km",milesToKm(10))
	fmt.Println("Kgs to Lbs",kgsToLbs(10))
	fmt.Println("Lbs to Kgs",lbsToKgs(10))
	fmt.Println("Celsius to Farhenheit",cToF(10))
	fmt.Println("Farhenheit to Celsius",fToC(10))
}