package main

import "fmt"

func main() {
	var n int
	fmt.Print("input: n = ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
