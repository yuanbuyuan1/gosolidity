package main

import (
	"fmt"
)

func main() {
	var a1 [5]int = [5]int{1, 2, 3, 4}

	fmt.Println(a1)
	a1[4] = 6 //使用数组下标对数组的第5个元素赋值
	fmt.Println(a1)
	s1 := [4]string{"lily", "lucy", "lilei"} //元素个数不能超过数组个数
	fmt.Println(s1)
}
