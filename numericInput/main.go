package main

import (
	"fmt"
	"sync"
	"time"
	"unicode"
)

type UserInput interface {
	Add(rune)
	GetValue() string
}

type NumericInput struct {
	input string
}

func (n *NumericInput) Add(d rune) {
	if unicode.IsDigit(d) {
		n.input += string(d)
	}
}

func (n *NumericInput) GetValue() string {
	return n.input
}

func main() {
	var input UserInput = &NumericInput{}

	input.Add('1')
	input.Add('n')
	input.Add('4')
	fmt.Println("str", input.GetValue())

	var wg sync.WaitGroup

	tasks := []func(){
		func() {
			time.Sleep(1 * time.Second)
			fmt.Println("Second 1, done")
		},
		func() {
			time.Sleep(2 * time.Second)
			fmt.Println("Task 2, done")
		},
		func() {
			fmt.Println("Done 1st")
		},
	}

	for _, task := range tasks {
		wg.Add(1)

		go func(t func()) {
			defer wg.Done()

			t()
		}(task)
	}

	wg.Wait()

}
