package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}

	return z
}

func main() {
	for i := 1.0; i <= 3; i++ {
		fmt.Println("calculating Sqrt of", i)
		fmt.Println(Sqrt(i))
	}
}
