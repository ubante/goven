package main

import "fmt"

type Human struct {
	name string
}

func (h *Human) talk() {
	fmt.Println(h.name, "says: blah blah blah")
}

type Android struct {
	Human
	model string
}

func main() {
	h := new(Human)
	h.name = "George"
	h.talk()

	a := new(Android)
	a.name = "R2"
	a.talk()

	c := Android{model: "golden"}
	c.name = "C3"
	c.talk()
}