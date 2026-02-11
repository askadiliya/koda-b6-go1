package main

import (
	"fmt"
	"math"
)

func keliling(r float64) float64 {
	return 2 * math.Pi * r
}

func luas(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	fmt.Println("Keliling:", keliling(14))
	fmt.Println("Luas:", luas(14))
}
