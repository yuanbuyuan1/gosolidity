package main

import (
	"fmt"
)

type Person struct {
	name  string
	age   int
	sex   string
	fight int
}

func (p *Person) setAge(age int) {
	p.age = age
}

type Superman struct {
	strength int
	speed    int
	p        Person
}

func (s *Superman) print() {
	fmt.Printf("%+v\n", s)
}

func main() {
	p1 := Person{"战五渣", 30, "man", 5}
	s1 := Superman{
		strength: 100000,
		speed:    19000,
		p:        p1, //注意逗号
	}
	s1.p.setAge(30)
	s1.print()

}
