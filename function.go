package main

import "fmt"

func add(x,y int) (int){
	return x + y
}

func multiply(x int, y int) (int){
	return x * y
}

func subtract(x, y int) (result int){
	result = x - y
	return
} 

func main()  {
	addition, multiplication, substraction :=  add(10,5), multiply(10,5), subtract(10,5) 
	fmt.Println(addition)
	fmt.Println(multiplication)
	fmt.Println(substraction)
}