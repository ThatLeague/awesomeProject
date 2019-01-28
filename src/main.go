package main

import "fmt"

func main () {
	println(Multiply(10,	20))
	fmt.Println(Multiply(20, -10));
	return
}


// Multiply numbers recusively without using Multiply opperator.
func Multiply(x int, y int) int {
	// The easiest and also the most important
	if y == 0 {
		return 0
	}

	// we can presume y is at least 1 so add it
	if y > 0 {
		return x + Multiply(x, y-1)
	}

	// if y is negative
	if y < 0 {
		return -Multiply(x, -y)
	}

	return -1
}
