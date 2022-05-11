package main

import (
	"fmt"
)

func add_sub(a, b int, f func(a, b int) int) int {
	return f(a, b)
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func main() {
	c := func(a, b int) int {
		return a + b
	}(10, 20)
	add_sub(10, 20, add) //10 + 20
	add_sub(10, 20, sub) //10 - 20
	fmt.Println(c)
}
