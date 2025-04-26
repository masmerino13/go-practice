package main

import (
	"fmt"
	"time"
)

func ReadPointer() {
	age := time.Now().Year() - 1989

	point := &age

	*point = 12

	fmt.Println(age, *point)
}
