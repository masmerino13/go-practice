package main

import (
	"strings"
)

func WordCount(s string) map[string]int {
	d := make(map[string]int)
	ss := strings.Fields(s)

	for _, v := range ss {
		d[v] = len(v)
	}

	return d
}
