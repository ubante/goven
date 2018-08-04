package main

import "fmt"

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

	dino1 := Dinosaur{"Deeno", 1}
	dino2 := Dinosaur{"Ross", 2}

	dinoPot := NewDInosaurReferenceManual()
	dinoPot.add(&dino1)
	dinoPot.add(&dino2)

	fmt.Println("\nThe Dino Pot:")
	fmt.Println(dinoPot)
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