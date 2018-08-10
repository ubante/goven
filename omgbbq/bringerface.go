package main

import "fmt"

/*
This is similar to blingerfaces.go

The secret ampersauce.
 */

type Bringer interface {
	getName() string
	setNext(b Bringer) error
	getNext() *Bringer
}

type GenericBringer struct {
	name string
	next *Bringer
}

func (gb GenericBringer) getName() string {
	return gb.name
}

func (gb GenericBringer) getNext() *Bringer {
	return gb.next
}

func (gb *GenericBringer) setNext(b Bringer) error {
	gb.next = &b
	fmt.Printf("Made %s the neighbor of %s.\n", b.getName(), gb.getName())

	//next := *gb.next
	next := *gb.getNext()
	fmt.Println("Proof:", next.getName())

	return nil
}

type AppleBringer struct {
	GenericBringer
}

type BananaBringer struct {
	GenericBringer
}

func validateBringerness(b Bringer) error {
	fmt.Println("Yes:", b)
	return nil
}

func main() {
	alice := AppleBringer{GenericBringer{name: "Alice"}}

	bert := BananaBringer{GenericBringer{name: "Bert"}}
	alice.setNext(&bert)
	fmt.Println(alice)
	fmt.Println(bert)

	fmt.Println("======================")
	gary := GenericBringer{name: "Gary"}
	bert.setNext(&gary)

	gertrude := GenericBringer{name: "Gertrude"}
	gary.setNext(&gertrude)
	fmt.Println(gary)
	fmt.Println(gertrude)
	fmt.Println("======================")

	validateBringerness(&alice)
	validateBringerness(&bert)
	validateBringerness(&gary)

	bringers := []Bringer{&alice, &bert, &gary}

	fmt.Println("======================")
	for _, bringer := range bringers {
		next := *bringer.getNext()
		fmt.Printf("%s -> %s\n", bringer.getName(), next.getName())
	}

}

/*
Made Bert the neighbor of Alice.
Proof: Bert
{{Alice 0xc0420541c0}}
{{Bert <nil>}}
======================
Made Gary the neighbor of Bert.
Proof: Gary
Made Gertrude the neighbor of Gary.
Proof: Gertrude
{Gary 0xc042054240}
{Gertrude <nil>}
======================
Yes: &{{Alice 0xc0420541c0}}
Yes: &{{Bert 0xc042054200}}
Yes: &{Gary 0xc042054240}
======================
Alice -> Bert
Bert -> Gary
Gary -> Gertrude

Process finished with exit code 0

 */