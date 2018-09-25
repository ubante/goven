package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"bufio"
	"strings"
	"strconv"
)

func read1() {
	fmt.Println("FIRST APPROACH")
	filename := "poker/matrix/holeCardValues_SklanskyMalmuth.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", filename)
		os.Exit(1)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println("Error stating file:", filename)
		os.Exit(2)
	}

	byteStat := make([]byte, stat.Size())
	_, err = file.Read(byteStat)
	if err != nil {
		fmt.Println("Error reading file:", filename)
		os.Exit(3)
	}

	str := string(byteStat)
	fmt.Println(str)
	fmt.Println()
}

func read2() {
	fmt.Println("SECOND APPROACH")
	filename := "poker/matrix/holeCardValues_SklanskyMalmuth.txt"
	byteSize, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error with ReadFile:", filename)
	}

	str := string(byteSize)
	fmt.Println(str)
	fmt.Println()
}

// https://stackoverflow.com/questions/5884154/read-text-file-into-string-array-and-write
func read3() {
	fmt.Println("THIRD APPROACH")
	filename := "poker/matrix/holeCardValues_SklanskyMalmuthModified.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", filename)
		os.Exit(1)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	ranks := []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}

	// https://stackoverflow.com/questions/44305617/nested-maps-in-golang
	var smMap = map[string]map[string]int{
		// How to do this better?
		"A": {},
		"K": {},
		"Q": {},
		"J": {},
		"T": {},
		"9": {},
		"8": {},
		"7": {},
		"6": {},
		"5": {},
		"4": {},
		"3": {},
		"2": {},
	}
	//fmt.Println("smMap:")
	//fmt.Println(smMap)
	//smMap := make(map[string]map[string]int)
	//smMap["K"]["K"] = 1

	for lineNumber, line := range lines {
		if lineNumber == 0 {
			continue  // Skip header line.
		}
		i := lineNumber - 1

		// Remove the hint before the colon and the colon itself.
		// https://programming.guide/go/split-string-into-slice.html
		colonSeparated := strings.Split(line, ":")

		// These are the Sklansky Malmuth values.
		commaSeparated := strings.Split(colonSeparated[1], ",")
		//fmt.Println("                    ", commaSeparated[2])

		// Assign these values to the map.
		for j := 0; j < len(commaSeparated); j++ {
			num, err := strconv.Atoi(strings.TrimSpace(commaSeparated[j]))
			if err != nil {
				fmt.Printf("Error converting value from string to int: [%s%s] -> %s\n", ranks[i], ranks[j],
					commaSeparated[j])
				os.Exit(3)
			}
			smMap[ranks[i]][ranks[j]] = num

		}
		fmt.Printf("%2d (%s): %s\n", i, ranks[i], line)
	}

	fmt.Println("This is [8A]:", smMap["8"]["A"])
	fmt.Println("This is [22]:", smMap["2"]["2"])
	fmt.Println("This is suited [J7]:  ", smMap["J"]["7"])
	fmt.Println("This is offsuite [7J]:", smMap["7"]["J"])


}

func main() {
	read1()
	read2()
	read3()
}

/*
FIRST APPROACH
//     A,  K,  Q,  J,  T,  9,  8,  7,  6,  5,  4,  3,  2
// A:  1,  1,  2,  2,  3,  5,  5,  5,  5,  5,  5,  5,  5
// K:  2,  1,  2,  3,  4,  6,  7,  7,  7,  7,  7,  7,  7
// Q:  3,  4,  1,  3,  4,  5,  7, 99, 99, 99, 99, 99, 99
// J:  4,  5,  5,  1,  3,  4,  6,  8, 99, 99, 99, 99, 99
// T:  6,  6,  6,  5,  2,  4,  5,  7, 99, 99, 99, 99, 99
// 9:  8,  8,  8,  7,  7,  3,  4,  5,  8, 99, 99, 99, 99
// 8: 99, 99, 99,  8,  8,  7,  4,  5,  6,  8, 99, 99, 99
// 7: 99, 99, 99, 99, 99, 99,  8,  5,  5,  6,  8, 99, 99
// 6: 99, 99, 99, 99, 99, 99, 99,  8,  6,  7,  7, 99, 99
// 5: 99, 99, 99, 99, 99, 99, 99, 99,  8,  6,  6,  7, 99
// 4: 99, 99, 99, 99, 99, 99, 99, 99, 99,  8,  7,  7,  8
// 3: 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99,  7,  8
// 2: 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99,  7

SECOND APPROACH
//     A,  K,  Q,  J,  T,  9,  8,  7,  6,  5,  4,  3,  2
// A:  1,  1,  2,  2,  3,  5,  5,  5,  5,  5,  5,  5,  5
// K:  2,  1,  2,  3,  4,  6,  7,  7,  7,  7,  7,  7,  7
// Q:  3,  4,  1,  3,  4,  5,  7, 99, 99, 99, 99, 99, 99
// J:  4,  5,  5,  1,  3,  4,  6,  8, 99, 99, 99, 99, 99
// T:  6,  6,  6,  5,  2,  4,  5,  7, 99, 99, 99, 99, 99
// 9:  8,  8,  8,  7,  7,  3,  4,  5,  8, 99, 99, 99, 99
// 8: 99, 99, 99,  8,  8,  7,  4,  5,  6,  8, 99, 99, 99
// 7: 99, 99, 99, 99, 99, 99,  8,  5,  5,  6,  8, 99, 99
// 6: 99, 99, 99, 99, 99, 99, 99,  8,  6,  7,  7, 99, 99
// 5: 99, 99, 99, 99, 99, 99, 99, 99,  8,  6,  6,  7, 99
// 4: 99, 99, 99, 99, 99, 99, 99, 99, 99,  8,  7,  7,  8
// 3: 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99,  7,  8
// 2: 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99,  7

THIRD APPROACH
 0 (A): // A:  1,  1,  2,  2,  3,  5,  5,  5,  5,  5,  5,  5,  5
 1 (K): // K:  2,  1,  2,  3,  4,  6,  7,  7,  7,  7,  7,  7,  7
 2 (Q): // Q:  3,  4,  1,  3,  4,  5,  7, 10, 30, 60, 99, 99, 99
 3 (J): // J:  4,  5,  5,  1,  3,  4,  6,  8, 15, 45, 99, 99, 99
 4 (T): // T:  6,  6,  6,  5,  2,  4,  5,  7, 20, 60, 99, 99, 99
 5 (9): // 9:  8,  8,  8,  7,  7,  3,  4,  5,  8, 15, 50, 99, 99
 6 (8): // 8: 11, 15, 15,  8,  8,  7,  4,  5,  6,  8, 20, 70, 99
 7 (7): // 7: 45, 60, 60, 20, 25, 20,  8,  5,  5,  6,  8, 25, 99
 8 (6): // 6: 99, 99, 99, 90, 80, 60, 18,  8,  6,  7,  7, 20, 99
 9 (5): // 5: 99, 99, 99, 99, 99, 99, 65, 20,  8,  6,  6,  7, 99
10 (4): // 4: 99, 99, 99, 99, 99, 99, 99, 90, 25,  8,  7,  7,  8
11 (3): // 3: 99, 99, 99, 99, 99, 99, 99, 99, 99, 40, 30,  7,  8
12 (2): // 2: 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 30,  7
This is [8A]: 11
This is [22]: 7
This is suited [J7]:   8
This is offsuite [7J]: 20

Process finished with exit code 0

 */