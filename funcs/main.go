package main

import "fmt"

func compute(fn func(int, int) int) int {
	return fn(5, 5)
}

func main() {
	do := func(x, y int) int {
		return x * y
	}

	fmt.Println(do(2, 2))

	fmt.Println(compute(do))
}
