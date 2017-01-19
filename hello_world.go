package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	// This should print out "Hello World 1 2 3"
	fmt.Printf("Hello %d %s %s %s\n", "World", 1, 2, 3)
}
