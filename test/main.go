package main

import (
	"fmt"
	"slices"
)

func uniqNames(a, b []string) []string {
	var res []string

	for _, str := range append(a, b...) {
		if !slices.Contains(res, str) {
			res = append(res, str)
		}
	}
	return res
}

func main() {
	fmt.Println(
		uniqNames(
			[]string{"Juan", "Pedro", "Mario", "Dan", "Chepe"},
			[]string{"Dan", "Juan"},
		),
	)
}
