package main

import (
	"fmt"
	"time"
	"os"
)

type Thing struct {
	name string
	pre *Thing
	post *Thing
}

type AllTheThings struct {
	things []*Thing
}

func (att *AllTheThings) addToEnd(t *Thing) {
	fmt.Println("Adding:", t.name)

	if len(att.things) == 0 {
		att.things = append(att.things, t)
		return
	}

	fmt.Println("Appending:", t.name)
	last := len(att.things)-1
	att.things[last].post = t
	t.pre = att.things[last]
	t.post = att.things[0]
	att.things[0].pre = t
	att.things = append(att.things, t)
}

func (att AllTheThings) printLinkList(reverse bool, t *Thing) {
	if t == nil {
		t = att.things[0]
	}
	thing := *t

	if reverse == false {
		if thing.post == att.things[0] {
			fmt.Println(thing.name)
			return
		}
	} else {
		if thing.pre == att.things[0] {
			fmt.Println(thing.name)
			return
		}
	}

	if reverse == false {
		fmt.Printf("%s -> ", thing.name)
	} else {
		fmt.Printf("%s <- ", thing.name)
	}
	time.Sleep(50 * time.Millisecond)

	if reverse == false {
		att.printLinkList(reverse, thing.post)
	} else {
		att.printLinkList(reverse, thing.pre)
	}

	return
}

// Position starts counting at 0 and ends at len(att.things)-1.
func (att *AllTheThings) insert(position int, thing *Thing) {
	length := len(att.things)
	if position > length {
		fmt.Println("You are trying to add to a position that is impossible.")
		fmt.Println("The slice length is", length, "so the last thing is at position", length-1)
		fmt.Println("so your position,", position, "is impossible.")
		os.Exit(9)
	}

	// If the position is at the front of the slice, do this.
	if position == 0 {
		fmt.Println("Adding to the front.")
		att.things[length-1].post = thing
		thing.pre = att.things[length-1]
		thing.post = att.things[0]
		att.things[0].pre = thing
		att.things = append([]*Thing{thing}, att.things...)
		return
	}

	// If the position is at the end of the slice, do this.
	if position == length {
		fmt.Println("Adding to the end.")
		att.things[length-1].post = thing
		thing.pre = att.things[length-1]
		thing.post = att.things[0]
		att.things[0].pre = thing
		att.things = append(att.things, thing)
		return
	}

	// Otherwise, do this.
	fmt.Println("Adding to position", position)
	att.things[position-1].post = thing
	thing.pre = att.things[position-1]
	thing.post = att.things[position]
	att.things[position].pre = thing

	att.things = append(att.things[:position], append([]*Thing{thing}, att.things[position:]...)...)
	return
}

func main() {
	all := AllTheThings{nil}

	t1 := Thing{"T1", nil, nil}
	all.addToEnd(&t1)
	t2 := Thing{"T2", nil, nil}
	all.addToEnd(&t2)
	t4 := Thing{"T4", nil, nil}
	all.addToEnd(&t4)
	t5 := Thing{"T5", nil, nil}
	all.addToEnd(&t5)

	all.printLinkList(false, nil)
	all.printLinkList(true, nil)

	t3 := Thing{"T3", nil, nil}
	all.insert(2, &t3)
	all.printLinkList(false, nil)
	all.printLinkList(true, nil)

	for i, thing := range all.things {
		//fmt.Println(i, thing.name)
		fmt.Printf("%d:              %s <- %s -> %s\n", i, thing.pre.name, thing.name, thing.post.name)
		fmt.Printf("%p <- %p -> %p\n", thing.pre, thing, thing.post)
	}

	fmt.Println("==================")
	a := Thing{"A", nil, nil}
	b := Thing{"B", &a, nil}
	a.post = &b

	d := Thing{"D", &b, nil}
	b.post = &d

	e := Thing{"E", &d, &a}
	d.post = &e
	a.pre = &e

	things := []*Thing{&a, &b, &d, &e}
	for i, thing := range things {
		fmt.Printf("%d:              %s <- %s -> %s\n", i, thing.pre.name, thing.name, thing.post.name)
		fmt.Printf("%p <- %p -> %p\n", thing.pre, thing, thing.post)
	}
	fmt.Println()

	c := Thing{"C", &b, &d}
	b.post = &c
	d.pre = &c

	things = append(things[:2], append([]*Thing{&c}, things[2:]...)...)
	for i, thing := range things {
		fmt.Printf("%d:              %s <- %s -> %s\n", i, thing.pre.name, thing.name, thing.post.name)
		fmt.Printf("%p <- %p -> %p\n", thing.pre, thing, thing.post)
	}
}