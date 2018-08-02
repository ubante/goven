package main

import "fmt"

type Animal struct {
	name string
}

// Note that the receiver isn't a pointer.
func (a Animal) String() string {
	return fmt.Sprintf(a.name)
}

// https://golang.org/pkg/fmt/
type X string

func (x X) String() string { return fmt.Sprintf("<%s>", string(x)) }

func main() {
	cat := Animal{"sylvester"}

	fmt.Println("The name is", cat.name)
	fmt.Println("And also", cat)

	xxx := X("xXx")
	fmt.Println("Triple", xxx)
}