package main

import (
	"fmt"
	"math"
)

const EPS float64 = 1e-15

// 5 iterations to obtain result
func Sqrt(x float64) float64 {
	z := float64(1)
	sqrt := math.Sqrt(x)
	for math.Abs(z - sqrt) > EPS {
		z -= (z * z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func Sqrt2(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		z -= (z * z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}