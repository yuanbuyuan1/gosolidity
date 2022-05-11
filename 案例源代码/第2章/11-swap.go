package main

import (
	"fmt"
)

func main() {
	a, b := 10, 20
	//一定要用a和b去接收返回值
	a, b = swap2(a, b)
	fmt.Println(a, b)
}

func swap2(a, b int) (int, int) {
	return b, a
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

// func swap(a, b int) {
// 	temp := a
// 	a = b
// 	b = temp
// }
