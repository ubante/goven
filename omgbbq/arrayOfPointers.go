package main

import (
	"fmt"
	"sort"
)

type SuperHero struct {
	factor0    int
	factor1    int
	factor2    int
	factor3    int
	allFactors [4]*int
}

func NewSuperHero(factors []int) *SuperHero {
	var sh SuperHero

	sh.factor0 = factors[0]
	sh.factor1 = factors[1]
	sh.factor2 = factors[2]
	sh.factor3 = factors[3]

	// How to do the below in one line?
	//sh.allFactors = []*int{&sh.factor0, &sh.factor1, &sh.factor2, &sh.factor3,}
	sh.allFactors[0] = &sh.factor0
	sh.allFactors[1] = &sh.factor1
	sh.allFactors[2] = &sh.factor2
	sh.allFactors[3] = &sh.factor3

	return &sh
}

func (sh SuperHero) String() string {
	returnedString := fmt.Sprintf("%d %d %d %d: ", sh.factor0, sh.factor1, sh.factor2, sh.factor3)

	for _, f := range sh.allFactors {
		returnedString += fmt.Sprintf("%d ", *f)
	}

	return returnedString
}

func (sh *SuperHero) incFactors() {
	sh.factor0++
	sh.factor1++
	sh.factor2++
	sh.factor3++
}

func (sh SuperHero) getOrderedFactors() []int {
	var orderedFactors []int
	for _, factor := range sh.allFactors {
		orderedFactors = append(orderedFactors, *factor)
	}
	sort.Ints(orderedFactors)

	return orderedFactors
}

func main() {
	moreThings := [5]int{5, 6, 7, 8, 9}
	fmt.Println("more things:", moreThings)

	batman := NewSuperHero([]int{4, 5, 3, 2})
	fmt.Println("batman:", batman)
	batman.incFactors()
	fmt.Println("batman:", batman)

	fmt.Println("ordered:", batman.getOrderedFactors())
}

/*
more things: [5 6 7 8 9]
batman: 4 5 3 2: 4 5 3 2
batman: 5 6 4 3: 5 6 4 3
ordered: [3 4 5 6]

*/
