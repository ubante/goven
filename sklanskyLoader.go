package main

import (
	"fmt"
	"goven/poker/matrix"
)

func main() {
	fmt.Println("This will read in a SM model no more than once.")

	score1 := matrix.GetScore()
	score2 := matrix.GetScore()
	score3 := matrix.GetScore()
	fmt.Println("1. The filename is:", score1.GetFilename())
	fmt.Println("2. The filename is:", score2.GetFilename())
	fmt.Println("3. The filename is:", score3.GetFilename())

	score1.SetFilename("new_filename")
	fmt.Println("1. The filename is:", score1.GetFilename())
	fmt.Println("2. The filename is:", score2.GetFilename())
	fmt.Println("3. The filename is:", score3.GetFilename())
}