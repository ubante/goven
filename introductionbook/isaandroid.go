package main

import "fmt"

// This creates a base class, subclass and interface.  And shows
// that you can create a slice of the interface.


type Humanoid struct {
	name string
	weight int
}

func (h *Humanoid) talk() {
	fmt.Println(h.name, "says: blah blah blah")
}

func (h Humanoid) eat(food string) {
	fmt.Println(h.name, "eats a", food)
}

func (h Humanoid) verb() {

}

type Android struct {
	Humanoid
	model string
}

func (a Android) eat(food string) {
	fmt.Println(a.name, "is an android.  Stop trying to feed it.")
}

type Eater interface {
	verb()

	// I think I can include params in method signature.  The below does.
	// https://medium.com/golangspec/interfaces-in-go-part-i-4ae53a97479c
	eat(food string)  // receiver can't be a pointer?  omgwtfbbq!

	// It is possible:
	// https://stackoverflow.com/questions/33936081/golang-method-with-pointer-receiver
	// https://stackoverflow.com/questions/45652560/interfaces-and-pointer-receivers
	//   https://play.golang.org/p/y7VITjHZ0-D
}

type Speaker interface {
	Speak()
}

func anonymizeName(h *Humanoid) {
	h.name = "new-random-name"
}

type Platoon struct {

}

func main() {
	h := new(Humanoid)
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

	fred := Android{
		Humanoid: Humanoid{
			name: "Fred",
			weight: 560,
		},
		model: "Bizarro",
	}
	fmt.Printf("\n%s weighs %d lb and is a %s model.\n\n", fred.name, fred.weight, fred.model)

	cyclonSquad := []Eater{
		Humanoid{
			name: "Edith",
			weight: 110,
		},
		Android{
			Humanoid: Humanoid{
				name: "Annie",
				weight: 623},
			model: "notacylon",
	    },
	}

	cyclonSquad = append(cyclonSquad, *h)
	cyclonSquad = append(cyclonSquad, *a)
	cyclonSquad = append(cyclonSquad, c)
	for i, humanoid := range cyclonSquad {
		fmt.Println(i+1, ":", humanoid)

		// The below is an error because the Eater interface does not have a name field.
		//fmt.Println(i, ":", humanoid.name)
		humanoid.eat("apples")
	}

}

/*
George says: blah blah blah
random says: blah blah blah
random eats a banana

R2 says: blah blah blah

C3 says: blah blah blah
C3 is an android.  Stop trying to feed it.

Fred weighs %!i(int=560) lb and is a Bizarro model.

0 : {Edith 110}
Edith eats a apples
1 : {{Annie 623} notacylon}
Annie is an android.  Stop trying to feed it.
 */