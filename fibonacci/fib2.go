package main

import "fmt"

func Fib2(n int) {
	result, a, b := 0, 0, 1

	for i := 0; i < n; i++ {
		result = a

		a, b = b, a+b

		fmt.Println("Fib num: ", result)
	}
}
