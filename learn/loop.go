package main

import "fmt"

func DoLoop() {
	count := 0

	for i := 0; i < 10; i++ {
		count += i
	}

	fmt.Println("Total is: ", count)
}

func OptionalOpt() {
	sum := 1

	for sum < 50 {
		fmt.Println(sum)
		sum += sum
	}

	fmt.Println(sum)
}
