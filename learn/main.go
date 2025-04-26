package main

import "fmt"

func sum(a, b int) (int, int) {
	return a, b + 2
}

func removeN(nums []int, val int) int {
	k := 0

	for _, v := range nums {
		if v != val {
			nums[k] = v
			k++
		}
	}

	return k
}

func main() {
	// fmt.Println("Welcome to playground!", math.Sqrt(7))
	// n1, n2 := Split(2)
	// fmt.Println("Today is", n1, n2)
	// fmt.Println(Split(10))
	// DoLoop()
	// OptionalOpt()

	fmt.Println(CheckIf("no"))
	fmt.Println(CheckIf("yes"))
	ReadPointer()
	ShowStructs()
	ShowMap()
	WordCount("hi there you hello")

	l1 := []int{1, 2, 3}
	a := 3
	fmt.Println(removeN(l1, a))
}
