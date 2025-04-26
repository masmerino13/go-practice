package main

import "fmt"

func SumInts(n map[string]int64) int64 {
	var t int64

	for _, v := range n {
		t += v
	}

	return t
}

func SumFloats(n []int64) int64 {
	var t int64

	for _, v := range n {
		t += v
	}

	return t
}

func GenSum()

func main() {
	n := map[string]int64{
		"one":  1,
		"one2": 3,
	}

	n2 := []int64{2, 4, 6}

	fmt.Println("total: ", SumInts(n))
	fmt.Println("total: ", SumFloats(n2))
}
