package main

import "fmt"

func simple_defer(){
	defer fmt.Println("First Line")

	fmt.Println("Last Line")
}

func increment_return_value() (i int) {
    defer func() { i++ }()
    return 1
}

func defer_in_loop(){
	fmt.Println("counting")

	for i := 0; i <= 5; i++ {
		defer fmt.Println(i) 
	}

	fmt.Println("done")
}

func main() {
	
	defer simple_defer() // defer in method call
	defer_in_loop()
	fmt.Println(increment_return_value())
	
}


// defer uses stack (LIFO)