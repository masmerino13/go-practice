package main

import "fmt"

type Vertex struct {
	x int
	y int
}

func ShowStructs() {
	v := Vertex{3, 5}
	p := &v
	p.x = 6

	fmt.Println(p.x)
}
