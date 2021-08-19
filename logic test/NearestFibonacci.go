package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
func main() {
	var n int
	fmt.Print("input: ")
	fmt.Scan(&n)

	fib(n)
	if n >= 0 {
		fmt.Println(fib(n))
	}

}
