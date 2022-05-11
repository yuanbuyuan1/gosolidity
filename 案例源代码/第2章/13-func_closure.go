//函数闭包
package main

import (
	"fmt"
)

func main() {
	nextnumber := getSequence() //nextnumber 是一个函数，可调用
	fmt.Println(nextnumber())
	fmt.Println(nextnumber())
	f := getSequence()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(nextnumber())
}

//序列函数
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
