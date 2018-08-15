package main

import "fmt"

/*
This is similar to singerfaces.go
 */

type Blinger interface {
	alsoPrint(s string)
}

type GenericBlinger struct {
	name string
}

func (gb *GenericBlinger) alsoPrint(s string) {
	fmt.Println("Nme is:", gb.name, "and heres:", s)
}

func printStuff(b Blinger) {
	b.alsoPrint("yes yes yes")
}

func main() {
	gail := GenericBlinger{name: "Gail"}
	printStuff(&gail)  // The ampersand is the secret sauce.
}
