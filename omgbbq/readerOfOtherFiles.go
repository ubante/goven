package main

import (
	"fmt"
	"goven/omgbbq/other"
)

// This type seems unreachable from other packages.
type TheThingInTheParentDir struct {
	MyGuts string
}

func main() {
	fmt.Println("It's lame that I can't do this.")

	sot := other.SomeOtherThing{}
	sot.InnardsCapitalized = "this is visible"
	fmt.Println("This is the other thing:", sot)  // works
	sot.CallBack()
}