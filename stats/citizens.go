package main

import "fmt"

// http://spf13.com/post/is-go-object-oriented/

type Citizens struct {
	Country string
	Person
}

type Person struct {
	Name string
	Address Address
	nextPerson *Person
	previousPerson *Person
}

func newPerson(name string) Person {
	anAddress := Address{"nonum", "nostreet"}
	aPerson := Person{name, anAddress, nil, nil}

	return aPerson
}

type Address struct {
	Number string
	Street string
}

// method
func (a Address) getNumber() string {
	return a.Number
}

type A struct{}

type B *A

func (a *A) Read() {}

// function
func Sqrt(x float64) float64 {
	z := 1.0
	//diff := 1000000.0
	//for i := 1; diff > 0.001; i++ {
	for i := 1; i <= 10; i++ {
		diff := (z*z - x) / (2*z)
		z -= diff
		// fmt.Println("z = ", z)
		fmt.Printf("%d: diff = %g, z = %g\n", i, diff, z)
	}

	return z
}

type Things struct {
	intList []int
	stringList []string
}

func main() {


	//var b B
	//b.Read()
	fmt.Println(Sqrt(2))
	fmt.Println(123456.3e-5)

	fmt.Println("======================================")
	var addr1 Address
	addr1.Number = "11"
	addr1.Street = "First St"
	fmt.Println("Number = ", addr1.getNumber())

	fmt.Println("======================================")
	addr2 := Address{"22", "Second St"}
	fmt.Println("Number = ", addr2.getNumber())

	fmt.Println("======================================")
	//var var list []int{1, 2}
	//list := []int{1, 2}
	var list []int
	fmt.Println(list)
	list = append(list, 33)
	fmt.Println(list)

	fmt.Println("======================================")
	//numbers := [20, 21, 22]
	numbers := []int{20, 21, 22}
	fmt.Println(numbers)

	//var things Things
	//things.intList := [4, 3, 4]


	fmt.Println("======================================")
	// https://blog.golang.org/slices
	// Create a couple of starter slices.
	slice := []int{1, 2, 3}
	slice2 := []int{55, 66, 77}
	fmt.Println("Start slice: ", slice)
	fmt.Println("Start slice2:", slice2)

	// Add an item to a slice.
	slice = append(slice, 4)
	fmt.Println("Add one item:", slice)

	// Add one slice to another.
	slice = append(slice, slice2...)
	fmt.Println("Add one slice:", slice)

	// Make a copy of a slice (of int).
	slice3 := append([]int(nil), slice...)
	fmt.Println("Copy a slice:", slice3)

	// Copy a slice to the end of itself.
	fmt.Println("Before append to self:", slice)
	slice = append(slice, slice...)
	fmt.Println("After append to self:", slice)

	//var s []byte
	//s = make([]byte, 5, 5)
	// s == []byte{0, 0, 0, 0, 0}


}