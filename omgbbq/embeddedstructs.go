package main

import "fmt"

type A struct {
	rootString string
}

type B struct {
	a1 A
	a2 A
}

type BBB struct {
	ptrA1 *A
	ptrA2 *A
}

type C struct {
	b1 B
	b2 B
}

type Lifeform struct {
	kingdom string
}

type Cat struct {
	Lifeform
	name string
}

func main() {

	bugger := Cat{name: "Sylvester"}
	bugger.kingdom = "Animale"
	fmt.Println("Cat's name is", bugger.name)
	fmt.Printf("%s is a %s\n", bugger.name, bugger.kingdom)

	aA := A{""}
	fmt.Println("A is", aA)

	pA := &aA
	fmt.Println("Pointer to an A is", pA)

	var s1 string
	s1 = "string 1"
	fmt.Println(s1)

	s2 := "string 2"
	fmt.Println(s2)

	var s3 string = "string 3"
	fmt.Println(s3)

	s4 := A{"string 4"}
	fmt.Println(s4.rootString)

	//var pC *C
	//pC = nil

	var ptrC *C = nil
	fmt.Println("Pointer to a C is", ptrC)

	var bbb *BBB
	bbb = nil
	fmt.Println("Pointer to a BBB is", bbb)
}
