package main

import "fmt"

type Human struct {
	name string
}

func (h *Human) talk() {
	fmt.Println(h.name, "says: blah blah blah")
}

func (h *Human) eat(food string) {
	fmt.Println(h.name, "eats a", food)
}

type Android struct {
	Human
	model string
}

func (a *Android) eat(food string) {
	fmt.Println(a.name, "is an android.  Stop trying to feed it.")
}

func anonymizeName(h *Human) {
	h.name = "random"
}

func main() {
	h := new(Human)
	h.name = "George"
	h.talk()
	anonymizeName(h)
	h.talk()
	h.eat("banana")

	fmt.Println()
	a := new(Android)
	a.name = "R2"
	a.talk()
	//anonymizeName(a)
	//a.talk()

	fmt.Println()
	c := Android{model: "golden"}
	c.name = "C3"
	c.talk()
	c.eat("carrot")
}