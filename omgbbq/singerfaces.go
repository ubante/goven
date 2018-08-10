package main

import (
	"fmt"
)

/*
This shows how to have subclasses.  The big lesson is to use pointers of
a subclasses when sending it to a function that expects the interface.

https://medium.com/@saiyerram/go-interfaces-pointers-4d1d98d5c9c6

 */

type Singer interface {
	start()
	stop()
	setNext(Singer)
	getNext() *Singer
	getName() string
}

type GenericSinger struct {
	name string
	octave int
	next *Singer
}

func (gs GenericSinger) deepSing() {
	fmt.Printf("(%s) La la la.\n", gs.name)
}

func (gs GenericSinger) sing() {
	gs.deepSing()
}

func (gs GenericSinger) start() {
	fmt.Printf("%s starts to sing.\n", gs.name)
	gs.sing()
}

func (gs GenericSinger) stop() {
	fmt.Printf("%s stops singing.\n", gs.name)
}

func (gs *GenericSinger) setNext(s Singer) {
	gs.next = &s
}

func (gs GenericSinger) getNext() *Singer {
	return gs.next
}

func (gs GenericSinger) getName() string {
	return gs.name
}

type SopranoSinger struct {
	GenericSinger
}

func (ss SopranoSinger) stop() {
	fmt.Println("Sopranos never stop singing.")
}

type TenorSinger struct {
	GenericSinger
}

// Note how this does not get called.
func (ts TenorSinger) deepSing() {
	fmt.Printf("(%s) Low low low.\n", ts.name)
}

func (ts TenorSinger) start() {
	fmt.Printf("%s starts to sing and the crowd starts to tear.\n", ts.name)
	ts.sing()
}

func main() {
	lea := SopranoSinger{GenericSinger{name: "Lea", octave: 4}}
	andrea := TenorSinger{GenericSinger{name: "Andrea", octave: 2}}
	lea.setNext(&andrea)
	theCond := GenericSinger{name: "Bun", octave: 0}
 	andrea.setNext(&theCond)
	theCond.setNext(&lea)

	singers := []Singer{&lea, &andrea, &theCond}

	for _, singer := range singers {
		singer.start()
		defer singer.stop()
	}

	fmt.Println("======================")
	for i, singer := range singers {
		next := *singer.getNext()
		fmt.Printf("%d: %s -> %s\n", i, singer.getName(), next.getName())

	}
	fmt.Println("======================")

}

/*
Lea starts to sing.
(Lea) La la la.
Andrea starts to sing and the crowd starts to tear.
(Andrea) La la la.
Bun starts to sing.
(Bun) La la la.
======================
0: Lea -> Andrea
1: Andrea -> Bun
2: Bun -> Lea
======================
Bun stops singing.
Andrea stops singing.
Sopranos never stop singing.
 */