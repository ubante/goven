package main

import (
	"fmt"
	"time"
)

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func mainfunc() {
	val := uint(6)

	var fact uint
	fact = factorial(val)

	fmt.Printf("%d factorial is %d\n\n", val, fact)
}

type Animal struct {
	name string
	nextAnimal *Animal
}

type AnimalQueue struct {
	queue []Animal
}

func (aq *AnimalQueue) add(name string) {
	newAnimal := Animal{name, nil}

	fmt.Println("Adding", name)
	if aq.queue == nil {
		newAnimal.nextAnimal = &newAnimal
		aq.queue = append(aq.queue, newAnimal)
		fmt.Printf("%s is the first in queue which now has %d length.\n", name, len(aq.queue))
		fmt.Println(newAnimal)
		return
	}

	fmt.Println("Queue has size:", len(aq.queue))
	aq.queue[len(aq.queue)-1].nextAnimal = &newAnimal
	aq.queue = append(aq.queue, newAnimal)

	fmt.Printf("Added %s to the queue which now has %d length.\n", name, len(aq.queue))
	//fmt.Println(newAnimal)
	for _, a := range aq.queue {
		fmt.Println(a)
		fmt.Print("-->")
		fmt.Println(a.nextAnimal)
	}
	fmt.Println("-------------------")
}

func (aq *AnimalQueue) printName(a *Animal) {
	fmt.Println("Trying:", a.name)
	if a == nil {
		a = &aq.queue[0]
		fmt.Println("Starting to print the names:")
		fmt.Println(a.name)

		time.Sleep(1000 * time.Millisecond)
		aq.printName(a.nextAnimal)
		return
	}

	fmt.Printf("This animal is %s and its next animal is %s\n", a.name, a.nextAnimal.name)
	aq.printName(a.nextAnimal)

	time.Sleep(1000 * time.Millisecond)

	return
}

func mainAnimal() {
	animals := AnimalQueue{nil}
	animals.add("aardvark")
	animals.add("bull")
	animals.add("cat")
	animals.add("dog")

	fmt.Println("The first element is:", animals.queue[0].name)
	fmt.Println("After it is:", animals.queue[0].nextAnimal.name)
	fmt.Println("The second element is:", animals.queue[1].name)
	fmt.Println("After it is:", animals.queue[1].nextAnimal.name)
	fmt.Println("The third element is:", animals.queue[2].name)
	//fmt.Println("AKA:", animals.queue[0].nextAnimal.nextAnimal.name)
	fmt.Println("After it is:", animals.queue[2].nextAnimal.name)

	fmt.Println("=====================================")
	fmt.Println(animals)
	fmt.Println("=====================================")
	fmt.Println("0:", animals.queue[0].name)
	fmt.Println("0:", animals.queue[0])
	fmt.Println("1:", animals.queue[0].nextAnimal.name)
	fmt.Println("1:", &animals.queue[1])
	fmt.Println("1:", animals.queue[0].nextAnimal)
	fmt.Println("2:", &animals.queue[2])
	//fmt.Println("2:", animals.queue[0].nextAnimal.nextAnimal.name)


	fmt.Println("=====================================")
	//animals.printName(&animals.queue[0])
}

func main() {
	mainfunc()
	mainAnimal()
}