package main

import (
	"fmt"
	"sort"
)

type Dinosaur struct {
	name string
	length int
}

type DinosaurReferenceManual struct {
	totalLength int
	farm map[*Dinosaur]int
}

func NewDInosaurReferenceManual() DinosaurReferenceManual {
	var drm DinosaurReferenceManual

	drm.totalLength = 0
	drm.farm = make(map[*Dinosaur]int)

	return drm
}

func (drm DinosaurReferenceManual) String() string {
	toString := fmt.Sprintf("========\nTotal length: %d\n", drm.totalLength)

	for dino, dinoLength := range(drm.farm) {
		toString += fmt.Sprintf("%s -> %d\n", dino.name, dinoLength)
	}

	toString += "========"
	return toString
}

func (drm *DinosaurReferenceManual) add(d *Dinosaur) {
	drm.totalLength += d.length
	drm.farm[d] = d.length
}

func (drm *DinosaurReferenceManual) invertMap() map[int][]*Dinosaur {
	inv := make(map[int][]*Dinosaur)

	for dino, length := range drm.farm {
		inv[length] = append(inv[length], dino)
	}

	return inv
}

func dinoMain() {
	dino1 := Dinosaur{"Deeno", 1}
	dino2 := Dinosaur{"Ross", 2}

	dinoPot := NewDInosaurReferenceManual()
	dinoPot.add(&dino1)
	dinoPot.add(&dino2)
	dinoPot.add(&Dinosaur{"Trex", 2})
	dinoPot.add(&Dinosaur{"Alysaur", 2})
	dinoPot.add(&Dinosaur{"Steg", 2})
	dinoPot.add(&Dinosaur{"Mega", 3})

	fmt.Println("\nThe Dino Pot:")
	fmt.Println(dinoPot)

	invertedMap := dinoPot.invertMap()
	fmt.Println(invertedMap)

	// https://stackoverflow.com/questions/23330781/sort-go-map-values-by-keys
	var sortedLengths []int
	for length := range invertedMap {
		sortedLengths = append(sortedLengths, length)
	}
	sort.Ints(sortedLengths)

	for _, length := range sortedLengths {
		fmt.Printf("%d: ", length)
		var names []string
		for _, element := range invertedMap[length] {
			names = append(names, element.name)
		}
		sort.Strings(names)
		for _, name := range names {
			fmt.Printf("%s, ", name)
		}
		fmt.Println()
	}
/*
1: Deeno,
2: Alysaur, Ross, Steg, Trex,
3: Mega,
 */

	// This is sorted by length but the names are not sorted.
	//for _, length := range sortedLengths {
	//	fmt.Printf("%d: ", length)
	//	for _, element := range invertedMap[length] {
	//		fmt.Printf("%s, ", element.name)
	//	}
	//	fmt.Println()
	//}
/*
1: Deeno,
2: Ross, Trex, Alysaur, Steg,
3: Mega,
 */

	//This is not sorted by length.
	//for length, list := range invertedMap {
	//	fmt.Printf("%d: ", length)
	//	for _, element := range list {
	//		fmt.Printf("%s, ", element.name)
	//	}
	//	fmt.Println()
	//}
/*
2: Trex, Alysaur, Steg, Ross,
3: Mega,
1: Deeno,
 */
}

// https://play.golang.org/p/U67R66Oab8r
func main() {

	// To create an empty map, use the builtin `make`:
	// `make(map[key-type]val-type)`.
	m := make(map[string]int)

	// Set key/value pairs using typical `name[key] = val`
	// syntax.
	m["k1"] = 7
	m["k2"] = 13

	// Printing a map with e.g. `fmt.Println` will show all of
	// its key/value pairs.
	fmt.Println("map:", m)

	// Get a value for a key with `name[key]`.
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// The builtin `len` returns the number of key/value
	// pairs when called on a map.
	fmt.Println("len:", len(m))

	// The builtin `delete` removes key/value pairs from
	// a map.
	delete(m, "k2")
	fmt.Println("map:", m)

	// The optional second return value when getting a
	// value from a map indicates if the key was present
	// in the map. This can be used to disambiguate
	// between missing keys and keys with zero values
	// like `0` or `""`. Here we didn't need the value
	// itself, so we ignored it with the _blank identifier_
	// `_`.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// You can also declare and initialize a new map in
	// the same line with this syntax.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	dinoMain()
}

/**
map: map[k1:7 k2:13]
v1:  7
len: 2
map: map[k1:7]
prs: false
map: map[foo:1 bar:2]

The Dino Pot:
========
Total length: 3
Deeno -> 1
Ross -> 2
========
 */