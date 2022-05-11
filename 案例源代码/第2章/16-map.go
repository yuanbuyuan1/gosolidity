package main

import "fmt"

func main() {
	countryCapitalMap := make(map[string]string)

	// map插入key - value对,各个国家对应的首都
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Roma"
	countryCapitalMap["China"] = "BeiJing"
	countryCapitalMap["India "] = "New Delhi"

	fmt.Println(countryCapitalMap)

}
