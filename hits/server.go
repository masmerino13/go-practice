package main

import "fmt"

type Key struct {
	Path, Country string
}

func Add(m map[string]map[string]int, path, country string) map[string]map[string]int {
	mm, ok := m[path]

	if !ok {
		mm = make(map[string]int)
		m[path] = mm
	}

	mm[country]++

	return m
}

func main() {
	hits := make(map[string]map[string]int)
	Add(hits, "/home", "sv")
	Add(hits, "/home", "sv")
	Add(hits, "/home", "us")

	hitsS := make(map[Key]int)

	hitsS[Key{"/contact", "au"}]++

	fmt.Println("hits", hits)
	fmt.Println("hits using struct", hitsS)

}
