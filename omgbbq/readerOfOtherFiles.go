package main

import (
	"fmt"
	"goven/omgbbq/other"
	//"goven/omgbbq/"
)

func main() {
	fmt.Println("It's lame that I can't do this.")

	sot := other.SomeOtherThing{}
	fmt.Println("This is the other thing:", sot)  // works

	sdt := SameDirThing{}
	fmt.Println("This is the local other thing:", sdt)
}