package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}


func assign_in_condition(n, p, lim float64)float64{
	 //pn is scoped in this condional block only

	if pn :=  math.Pow(n, p); pn < lim {
		return pn
	} else {
		fmt.Printf("%g is not greater than %g\n", pn, lim)
	}
	return lim
}

func main() {
	fmt.Println(sqrt(16))
	fmt.Println(
		assign_in_condition(3, 2, 10),
		assign_in_condition(3, 3, 20),
	)
}