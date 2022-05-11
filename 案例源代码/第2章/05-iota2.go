//枚举
package main

import (
	"fmt"
)

const (
	login = iota // itoa = 0
	logout
	user    = iota + 1
	account = iota + 3
)

const (
	apple, banana = iota + 1, iota + 2
	peach, pear
	orange, mango
)

func main() {
	fmt.Println(login, logout, user, account)
	fmt.Println(apple, banana, peach, pear, orange, mango)
}
