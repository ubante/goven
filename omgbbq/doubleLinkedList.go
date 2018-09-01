package main

import "fmt"

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