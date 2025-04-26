package main

import "fmt"

func doCalc(n []int) int {
	n1, n2 := 0, 0

	for _, i := range n {
		if i > n1 {
			n2 = n1
			n1 = i
		}
	}

	fmt.Println(n1, n2)

	return n1 + n2
}

func main() {
	n := []int{2, 5, 7, 9, 1, 4}

	fmt.Println(doCalc(n))
}
