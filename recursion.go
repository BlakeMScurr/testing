package main

import (
	"fmt"
	"os"
	// "strconv"
)

func fib(n int) int {
	if n <= 1 {
		return 1
	}
	return fib(n-2) + fib(n-1)
}

func main() {
	// Sequence is 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377
	// Should print 5040
	input, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.println(fib(input))
}
