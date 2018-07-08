package main

import "fmt"

type athing struct {
	firstThing  int
	secondThing string
	lastThing
}

type lastThing struct {
	name string
}

func (a *athing) doSomething(thething string) {
	fmt.Printf("\nGone done %s.\n", thething)
	fmt.Printf("Now at %d weight.\n", a.firstThing)
}

func (a *athing) doubleDown() {
	a.firstThing *= 2
}

func main() {
	thisthing := athing{firstThing: 1, secondThing: "yup"}
	thisthing.doubleDown()
	thisthing.doubleDown()
	thisthing.doubleDown()
	thisthing.doSomething("thingthathingthingthing")

	canidothis := thisthing
	canidothis.doSomething("it")
	canidothis.doubleDown()
	canidothis.doSomething("it.  Oops")

	thisthing.doSomething("thingthathingthingthing")
}

