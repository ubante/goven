package main

import (
	"fmt"
	"apples"
	"pinterest.com/ads/services/codfish/omgbbqdonotcommitthis/apples"
)

type Fruit struct {}

func (f *Fruit) sauce() {
	fmt.Println("tasty object")
}

func main() {
	fmt.Println("What happens if you call pkg.Func when you have both:")
	fmt.Println(" (1) struct named pkg with method named Func")
	fmt.Println(" (2) pacakge named pkg with method named Func")
	fmt.Println()

	// (1)
	apple := new(Fruit)
	apple.sauce()

	// (2)
	// If we rename the applepkg package to apple, then the below would be
	// shadowed by the apple variable.
	applepkg.Sauce()

}

/*
What happens if you call pkg.Func when you have both:
(1) struct named pkg with method named Func
(2) pacakge named pkg with method named Func

tasty object
tasty package
 */