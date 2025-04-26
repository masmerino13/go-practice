package main

func fib() func() int {
	a, b := 0, 1

	return func() int {
		result := a
		a, b = b, a+b

		return result
	}
}

func main() {
	/*
		f := fib()

		for i := 0; i < 7; i++ {
			fmt.Println(f())
		}
	*/

	Fib2(10)
}
