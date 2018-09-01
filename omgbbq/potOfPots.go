package main

import "fmt"

type SuperPot struct {

}

type Fool struct {
	name string
	stack int
	bet int
	handStrength int
}

func (f *Fool) raise(amount int) {
	f.stack -= amount
	f.bet += amount
}

func main() {
	a := Fool{name: "AAA", stack: 1000}
	b := Fool{name: "BBB", stack: 500}
	c := Fool{name: "CCC", stack: 1000}

	players := []*Fool{&a, &b, &c}
	for _, p := range players {
		fmt.Println(p)
		p.raise(p.stack)
		fmt.Println(p)
		fmt.Println("----------")
	}

	// Let the allIn player have the strongest hand.
	b.handStrength = 1
	c.handStrength = 2
	a.handStrength = 3

	// Testing the flattening.
	// https://www.tutorialspoint.com/go/go_multi_dimensional_arrays.htm
	loopData := [3][6]int{
		{7, 11, 6, 0, 0, 0},
		{3, 12, 6, 5, 4, 0},
		{9, 13, 11, 3, 2, 8},
	}

	var score int
	for i, inner := range loopData {
		fmt.Printf("%d: %v\n", i, inner)
		score = 0
		for _, innerInner := range inner {
			score *= 100
			score += innerInner
			fmt.Println(score)
		}

		fmt.Printf("%d: %v %d\n", i, inner, score)
	}


}