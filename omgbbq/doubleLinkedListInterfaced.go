package main

import (
	"fmt"
	"time"
	"os"
)

type Thanger interface {
	addToEnd(t *Thang)
	getName() string
	getPre() *Thang
	getPost() *Thang
}

type Thang struct {
	name string
	pre *Thang
	post *Thang
}

type SubThang struct {
	Thang
}

func (t Thang) getName() string {
	return t.name
}

func (t Thang) getPre() *Thang {
	return t.pre
}

func (t Thang) getPost() *Thang {
	return t.post
}

type AllTheThangs struct {
	thangs []*Thanger
}

func (att *AllTheThangs) addToEnd(t *Thanger) {
	thang := *t
	fmt.Println("Adding:", thang.getName())

	if len(att.thangs) == 0 {
		att.thangs = append(att.thangs, t)
		return
	}

	fmt.Println("Appending:", thang.getName())
	last := len(att.thangs)-1
	att.thangs[last].post = t
	t.pre = att.thangs[last]
	t.post = att.thangs[0]
	att.thangs[0].pre = t
	att.thangs = append(att.thangs, t)
}

func (att AllTheThangs) printLinkList(reverse bool, t *Thang) {
	if t == nil {
		t = att.thangs[0]
	}
	thang := *t

	if reverse == false {
		if thang.post == att.thangs[0] {
			fmt.Println(thang.name)
			return
		}
	} else {
		if thang.pre == att.thangs[0] {
			fmt.Println(thang.name)
			return
		}
	}

	if reverse == false {
		fmt.Printf("%s -> ", thang.name)
	} else {
		fmt.Printf("%s <- ", thang.name)
	}
	time.Sleep(50 * time.Millisecond)

	if reverse == false {
		att.printLinkList(reverse, thang.post)
	} else {
		att.printLinkList(reverse, thang.pre)
	}

	return
}

// Position starts counting at 0 and ends at len(att.thangs)-1.
func (att *AllTheThangs) insert(position int, thang *Thang) {
	length := len(att.thangs)
	if position > length {
		fmt.Println("You are trying to add to a position that is impossible.")
		fmt.Println("The slice length is", length, "so the last thang is at position", length-1)
		fmt.Println("so your position,", position, "is impossible.")
		os.Exit(9)
	}

	// If the position is at the front of the slice, do this.
	if position == 0 {
		fmt.Println("Adding to the front.")
		att.thangs[length-1].post = thang
		thang.pre = att.thangs[length-1]
		thang.post = att.thangs[0]
		att.thangs[0].pre = thang
		att.thangs = append([]*Thang{thang}, att.thangs...)
		return
	}

	// If the position is at the end of the slice, do this.
	if position == length {
		fmt.Println("Adding to the end.")
		att.thangs[length-1].post = thang
		thang.pre = att.thangs[length-1]
		thang.post = att.thangs[0]
		att.thangs[0].pre = thang
		att.thangs = append(att.thangs, thang)
		return
	}

	// Otherwise, do this.
	fmt.Println("Adding to position", position)
	att.thangs[position-1].post = thang
	thang.pre = att.thangs[position-1]
	thang.post = att.thangs[position]
	att.thangs[position].pre = thang

	att.thangs = append(att.thangs[:position], append([]*Thang{thang}, att.thangs[position:]...)...)
	return
}

func main() {
	all := AllTheThangs{nil}

	t1 := Thang{"T1", nil, nil}
	all.addToEnd(&t1)
	t2 := Thang{"T2", nil, nil}
	all.addToEnd(&t2)
	t4 := Thang{"T4", nil, nil}
	all.addToEnd(&t4)
	t5 := Thang{"T5", nil, nil}
	all.addToEnd(&t5)

	all.printLinkList(false, nil)
	all.printLinkList(true, nil)

	t3 := Thang{"T3", nil, nil}
	all.insert(2, &t3)
	all.printLinkList(false, nil)
	all.printLinkList(true, nil)

	for i, thang := range all.thangs {
		//fmt.Println(i, thang.name)
		fmt.Printf("%d:              %s <- %s -> %s\n", i, thang.pre.name, thang.name, thang.post.name)
		fmt.Printf("%p <- %p -> %p\n", thang.pre, thang, thang.post)

/*
Adding: T1
Adding: T2
Appending: T2
Adding: T4
Appending: T4
Adding: T5
Appending: T5
T1 -> T2 -> T4 -> T5
T1 <- T5 <- T4 <- T2
Adding to position 2
T1 -> T2 -> T3 -> T4 -> T5
T1 <- T5 <- T4 <- T3 <- T2
0:              T5 <- T1 -> T2
0xc0420024c0 <- 0xc042002420 -> 0xc042002460
1:              T1 <- T2 -> T3
0xc042002420 <- 0xc042002460 -> 0xc0420024e0
2:              T2 <- T3 -> T4
0xc042002460 <- 0xc0420024e0 -> 0xc042002480
3:              T3 <- T4 -> T5
0xc0420024e0 <- 0xc042002480 -> 0xc0420024c0
4:              T4 <- T5 -> T1
0xc042002480 <- 0xc0420024c0 -> 0xc042002420
 */
	}

	//subAll := AllTheThangs{nil}
	//s1 := SubThang{Thang{"S1", nil, nil}}
	//subAll.addToEnd(s1)
	//s2 := SubThang{Thang{"S2", nil, nil}}
	//subAll.addToEnd(s2)

	fmt.Println("==================")
	a := Thang{"A", nil, nil}
	b := Thang{"B", &a, nil}
	a.post = &b

	d := Thang{"D", &b, nil}
	b.post = &d

	e := Thang{"E", &d, &a}
	d.post = &e
	a.pre = &e

	thangs := []*Thang{&a, &b, &d, &e}
	for i, thang := range thangs {
		fmt.Printf("%d:              %s <- %s -> %s\n", i, thang.pre.name, thang.name, thang.post.name)
		fmt.Printf("%p <- %p -> %p\n", thang.pre, thang, thang.post)
	}
	fmt.Println()

	c := Thang{"C", &b, &d}
	b.post = &c
	d.pre = &c

	thangs = append(thangs[:2], append([]*Thang{&c}, thangs[2:]...)...)
	for i, thang := range thangs {
		fmt.Printf("%d:              %s <- %s -> %s\n", i, thang.pre.name, thang.name, thang.post.name)
		fmt.Printf("%p <- %p -> %p\n", thang.pre, thang, thang.post)
	}
}