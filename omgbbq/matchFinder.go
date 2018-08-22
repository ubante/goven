package main

import "fmt"

func main() {
	frequency := make(map[int]int)
	frequency[3] = 2
	frequency[5] = 2
	frequency[14] = 1

	matches := make(map[int][]int)
	for rank := range frequency {
		matches[frequency[rank]] = append(matches[frequency[rank]], rank)
	}

	fmt.Println(matches)  // map[1:[14] 2:[3 5]]
}
