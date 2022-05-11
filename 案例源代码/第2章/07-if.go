package main

import (
	"fmt"
)

func main() {
	var a = 10
	if a > 10 {
		fmt.Println("a bigger than 10.")
	} else if a < 10 {
		fmt.Println("a less than 10.")
	} else {
		fmt.Println("a equal 10.")
	}
}
