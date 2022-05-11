package main

import (
	"fmt"
)

const (
	Sun = iota
	Mon = iota
	Tue
	Wed
	Thu
	Fri
	Sat
)

func main() {
	fmt.Println(Sun, Mon, Tue, Wed, Thu, Fri, Sat)
}
