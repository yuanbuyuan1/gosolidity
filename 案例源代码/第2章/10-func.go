package main

import (
	"fmt"
)

func main() {
	a, b := 10, 20
	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
}

func add(a int, b int) int {
	return a + b
}
