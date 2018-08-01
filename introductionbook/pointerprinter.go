package main

import "fmt"

type Dude struct {
	name string
	number int
}

func (d *Dude) incNumber() {
	d.number += 1
}

type Car struct {
	name string
	members []*Dude
	driver *Dude
}

func (c *Car) incFirstMemberDudeNumber() {
	//c.members[0].number += 1  // This works.

	c.members[0].incNumber()  // This works.
}

func (c *Car) incDriver() {
	originalNumber := c.driver.number
	c.driver.number += 1
	fmt.Printf("%s had a number of %d but now it is %d.\n", c.driver.name, originalNumber, c.driver.number)
}

func main() {
	name := "Phyllis"
	pname := &name
	pname2 := &name

	fmt.Println(name, pname, pname2)

	d := Dude{name, 100}
	fmt.Println(d.name, "has number", d.number)
	d.incNumber()
	fmt.Println(d.name, "has number", d.number)

	//c := Car{"Camaro", nil, nil}
	var c Car
	c.name = "Camaro"
	c.members = append(c.members, &d)
	c.driver = &d

	fmt.Println(c.members[0].name, "has these numbers:", c.members[0].number)
	c.incFirstMemberDudeNumber()
	// c    .incFirstMemberDudeNumber()
	// table.postBlinds()

	// c.members[0]    .incNumber()
	// t.bigBlindPlayer.payBlind(t.bigBlindValue)

	fmt.Println(c.members[0].name, "has these numbers:", c.members[0].number)
	c.incFirstMemberDudeNumber()
	fmt.Println(c.members[0].name, "has these numbers:", c.members[0].number)
	c.incDriver()
	fmt.Println(c.members[0].name, "has these numbers:", c.members[0].number)
	c.incDriver()
	fmt.Println(c.members[0].name, "has these numbers:", c.members[0].number)
	c.incDriver()

	dp1 := &d
	dp2 := dp1
	fmt.Println("dp1 is", dp1.name)
	fmt.Println("dp2 is", dp2.name)

	d2 := Dude{"Amaryllis", 80}
	dp2 = &d2
	fmt.Println("dp1 is", dp1.name)
	fmt.Println("dp2 is", dp2.name)


}
