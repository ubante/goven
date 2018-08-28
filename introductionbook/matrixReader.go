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
	filename := "poker/holeCardValues_SklanskyMalmuth.txt"
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
	filename := "poker/holeCardValues_SklanskyMalmuth.txt"
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
	//filename := "poker/holeCardValues_SklanskyMalmuth.txt"
	filename := "poker/holeCardValues_SklanskyMalmuthModified.txt"
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
		"A": map[string]int{},
		"K": map[string]int{},
		"Q": map[string]int{},
		"J": map[string]int{},
		"T": map[string]int{},
		"9": map[string]int{},
		"8": map[string]int{},
		"7": map[string]int{},
		"6": map[string]int{},
		"5": map[string]int{},
		"4": map[string]int{},
		"3": map[string]int{},
		"2": map[string]int{},
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