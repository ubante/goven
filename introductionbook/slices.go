package main

import "fmt"

func main() {
	var aslice []int

	fmt.Println("a:", aslice, len(aslice), cap(aslice))

	aslice = append(aslice, 11)
	fmt.Println("a:", aslice, len(aslice), cap(aslice))

	bslice := make([]int, 0)
	fmt.Println("b:", bslice, len(bslice), cap(bslice))

	bslice = append(bslice, 21)
	fmt.Println("b:", bslice, len(bslice), cap(bslice))

}