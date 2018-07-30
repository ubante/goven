package main

import "fmt"

type Person struct {
	name string
	next *Person
}

type Line struct {
	people []Person
}

type PointerLine struct {
	people []*Person  // the learnings
}

func main() {
	var line Line
	var pline PointerLine

	a := Person{"Adam", nil}
	fmt.Println(a)
	line.people = append(line.people, a)
	fmt.Println(line)
	fmt.Println("===")
	pline.people = append(pline.people, &a)
	fmt.Println(pline)
	fmt.Println("-----------")

	b := Person{"Beth", nil}
	fmt.Println(b)
	line.people = append(line.people, b)
	fmt.Println(line)
	line.people[0].next = &b
	fmt.Println(line)
	fmt.Println("===")
	pline.people[0].next = &b
	pline.people = append(pline.people, &b)
	fmt.Println(pline)
	fmt.Println("-----------")

	c := Person{"Carl", nil}
	fmt.Println(c)
	line.people = append(line.people, c)
	fmt.Println(line)
	line.people[1].next = &c
	fmt.Println(line)
	fmt.Println("===")
	pline.people[1].next = &c
	pline.people = append(pline.people, &c)
	fmt.Println(pline)
	fmt.Println("-----------")

	fmt.Print(line.people[0].name)
	fmt.Print("->")
	fmt.Print(line.people[0].next.name)
	fmt.Print("->")
	fmt.Print(line.people[0].next.next)

	fmt.Println("")
	fmt.Println(b)
	fmt.Println(&b)
	fmt.Println(line.people[1])

	//fmt.Print(line.people[0].next.next.name)
	//fmt.Print("->")
	//fmt.Println(line.people[0].next.next.next)
	fmt.Println("===")
	fmt.Print(pline.people[0])
	fmt.Print("->")
	fmt.Print(pline.people[1])
	fmt.Print("->")
	fmt.Print(pline.people[2])
	fmt.Println()

	fmt.Println(pline.people[0])
	fmt.Println(pline.people[0].next)
	fmt.Println(pline.people[0].next.next)


	fmt.Println("----------------------")

	f := Person{"Fred", nil}
	e := Person{"Eric", &f}
	d := Person{"Doug", &e}

	fmt.Println(d, e, f)
	betterLine := Line{}
	betterLine.people = append(betterLine.people, d)
	fmt.Println(betterLine)
	betterLine.people = append(betterLine.people, e)
	fmt.Println(betterLine)
	betterLine.people = append(betterLine.people, f)
	fmt.Println(betterLine)

	fmt.Print(betterLine.people[0])
	fmt.Print("->")
	fmt.Print(betterLine.people[0].next)
	fmt.Print("->")
	fmt.Print(betterLine.people[0].next.next)

}