package main

import "fmt"
import "math/cmplx"

var c, python, java bool
const Pi = 3.14

var(
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func basic_variables (){
	var count int
	var golang = "Hello"
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	fmt.Println(count, c, python, java, golang)
}


func print_data_types_with_value(){
	fmt.Printf("Type: %T Value: %v\n", Pi, Pi)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

func main(){
	basic_variables()
	print_data_types_with_value()

}