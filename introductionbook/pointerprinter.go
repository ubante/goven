package main

import "fmt"

func main() {
	name := "Phyllis"
	pname := &name
	pname2 := &name

	fmt.Println(name, pname, pname2)


}
