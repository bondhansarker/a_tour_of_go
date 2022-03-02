package main

import "fmt"

func infinite_loop(){
	for {
	}
}

func for_loop(){
	var sum int
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func for_behaving_like_while(){
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
}

func for_loop_without_initialization_and_post_statement(){
		sum := 1
		for ; sum < 100; {
			sum += sum
		}
		fmt.Println(sum)
}

func main() {
	for_loop()
	for_loop_without_initialization_and_post_statement()
	for_behaving_like_while()
	// infinite_loop()

}