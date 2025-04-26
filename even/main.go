package main

import "fmt"

type person struct {
	fName string
	lName string
}

func main() {
	alex := person{
		fName: "Ale",
		lName: "May",
	}

	var jhon person

	fmt.Println(alex)

	fmt.Printf("%+v", jhon)
}
