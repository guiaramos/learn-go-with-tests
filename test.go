package main

import "fmt"

type person struct {
	first string
}

type human interface {
	speak()
}

func (p person) speak() {
	fmt.Printf("Hello, I am " + p.first)
}

func talk(h human) {
	h.speak()
}

func main() {
	p := person{first: "gui"}
	talk(p)
}
