package main

func Fib2(n int, c chan int) {
	a, b := 0, 1

	for i := 0; i < n; i++ {
		c <- a
		a, b = b, a+b
	}

	close(c)
}
