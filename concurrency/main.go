package main

import (
	"fmt"
	"sync"
)

func proccessData(wg *sync.WaitGroup, result *int, v int) {
	defer wg.Done()

	nv := v * 2

	fmt.Println("r", result)

	*result = nv
}

func main() {
	input := []int{2, 3, 4, 5}
	result := make([]int, len(input))
	var wg sync.WaitGroup

	for i, v := range input {
		wg.Add(1)

		go proccessData(&wg, &result[i], v)
	}

	wg.Wait()

	fmt.Println(result)

	/*
		f, err := os.Create("aja.txt")

		if err != nil {
			panic(err)
		}

		final := 16777215

		for i := 0; i < final; i++ {
			_, err = f.WriteString(fmt.Sprintf("%06x\n", i))

			if err != nil {
				panic(err)
			}
		}
	*/

	/*
		c := make(chan int, 5)

		go Fib2(cap(c), c)

		for i := range c {
			fmt.Println(i)
		}
	*/
}
