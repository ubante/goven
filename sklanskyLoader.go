package main

import (
	"fmt"
	"goven/poker/matrix"
)

func main() {
	fmt.Println("This will read in a SM model no more than once.")

	score1 := matrix.GetSMScore()
	score2 := matrix.GetSMScore()
	score3 := matrix.GetSMScore()
	fmt.Println("1. The filename is:", score1.GetFilename())
	fmt.Println("2. The filename is:", score2.GetFilename())
	fmt.Println("3. The filename is:", score3.GetFilename())

	score1.SetFilename("new_filename")
	fmt.Println("1. The filename is:", score1.GetFilename())
	fmt.Println("2. The filename is:", score2.GetFilename())
	fmt.Println("3. The filename is:", score3.GetFilename())

	fmt.Println("1. This is AJ:", score1.GetScoreOfString("AJ"))
	fmt.Println("1. This is J7:", score1.GetScoreOfString("J7"))
	fmt.Println("2. This is J7:", score1.GetScoreOfString("J7"))

	fmt.Println("3. This is SA and HQ:", score3.GetScoreOfHoleCardStrings("SA", "HQ"))
	fmt.Println("3. This is HQ and SA:", score3.GetScoreOfHoleCardStrings("HQ", "SA"))
	fmt.Println("3. This is HQ and HA:", score3.GetScoreOfHoleCardStrings("HQ", "HA"))
	fmt.Println("3. This is HQ and SQ:", score3.GetScoreOfHoleCardStrings("HQ", "SQ"))
	fmt.Println("3. This is CK and C4:", score3.GetScoreOfHoleCardStrings("CK", "C4"))
	fmt.Println("3. This is C4 and CK:", score3.GetScoreOfHoleCardStrings("C4", "CK"))
}