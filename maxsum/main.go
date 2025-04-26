package main

import (
	"fmt"
	"math"
)

func findMaxSum(numbers []int) int {
	fMax, sMax := math.MinInt64, math.MinInt64

	for _, n := range numbers {
		if n > fMax {
			sMax = fMax
			fMax = n
		} else if n > sMax {
			sMax = n
		}
	}

	return fMax + sMax
}

func main() {
	list := []int{5, 9, 7, 11}
	fmt.Println(findMaxSum(list))
}
