package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) younger() int {
	return p.age - 1
}

func main() {
	fmt.Print("hello")
	a := 1
	bingo := 2
	c := a + bingo
	fmt.Println(c)
	d := 'A'
	fmt.Println(d)

	p := person{"Jim", "Dibb", 49}
	fmt.Println(p)
	fmt.Println(p.younger())

}
