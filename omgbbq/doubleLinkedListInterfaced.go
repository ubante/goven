package main

import (
	"fmt"
	"os"
	"time"
)

type Thanger interface {
	getName() string
	getPre() *Thanger
	setPre(preThang *Thanger)
	getPost() *Thanger
	setPost(postThang *Thanger)
}

type Thang struct {
	name string
	pre *Thanger
	post *Thanger
}

type SubThang struct {
	Thang
}

func (t Thang) getName() string {
	return t.name
}

func (t Thang) getPre() *Thanger {
	return t.pre
}

func (t *Thang) setPre(preThang *Thanger) {
	t.pre = preThang
}

func (t Thang) getPost() *Thanger {
	return t.post
}

func (t *Thang) setPost(postThang *Thanger) {
	t.post = postThang
}

type AllTheThangs struct {
	thangs []*Thanger
}

func (att *AllTheThangs) addToEnd(thang Thanger) {
	fmt.Println("Adding:", thang.getName())

	if len(att.thangs) == 0 {
		att.thangs = append(att.thangs, &thang)
		return
	}

	fmt.Println("Appending:", thang.getName())
	indexLast := len(att.thangs)-1
	lastThang := *att.thangs[indexLast]
	lastThang.setPost(&thang)
	thang.setPre(&lastThang)

	indexFirst := 0
	firstThang := *att.thangs[indexFirst]
	firstThang.setPre(&thang)
	thang.setPost(&firstThang)

	att.thangs = append(att.thangs, &thang)
}

func (att AllTheThangs) printLinkList(reverse bool, t *Thanger) {
	if t == nil {
		t = att.thangs[0]
	}
	thang := *t

	if reverse == false {
		if thang.getPost() == att.thangs[0] {
			fmt.Println(thang.getName())
			return
		}
	} else {
		if thang.getPre() == att.thangs[0] {
			fmt.Println(thang.getName())
			return
		}
	}

	if reverse == false {
		fmt.Printf("%s -> ", thang.getName())
	} else {
		fmt.Printf("%s <- ", thang.getName())
	}
	time.Sleep(50 * time.Millisecond)

	if reverse == false {
		att.printLinkList(reverse, thang.getPost())
	} else {
		att.printLinkList(reverse, thang.getPre())
	}

	return
}

// Position starts counting at 0 and ends at len(att.thangs)-1.
//func (att *AllTheThangs) insert(position int, thang *Thang) {
//	length := len(att.thangs)
//	if position > length {
//		fmt.Println("You are trying to add to a position that is impossible.")
//		fmt.Println("The slice length is", length, "so the last thang is at position", length-1)
//		fmt.Println("so your position,", position, "is impossible.")
//		os.Exit(9)
//	}
//
//	// If the position is at the front of the slice, do this.
//	if position == 0 {
//		fmt.Println("Adding to the front.")
//		att.thangs[length-1].post = thang
//		thang.pre = att.thangs[length-1]
//		thang.post = att.thangs[0]
//		att.thangs[0].pre = thang
//		att.thangs = append([]*Thang{thang}, att.thangs...)
//		return
//	}
//
//	// If the position is at the end of the slice, do this.
//	if position == length {
//		fmt.Println("Adding to the end.")
//		att.thangs[length-1].post = thang
//		thang.pre = att.thangs[length-1]
//		thang.post = att.thangs[0]
//		att.thangs[0].pre = thang
//		att.thangs = append(att.thangs, thang)
//		return
//	}
//
//	// Otherwise, do this.
//	fmt.Println("Adding to position", position)
//	att.thangs[position-1].post = thang
//	thang.pre = att.thangs[position-1]
//	thang.post = att.thangs[position]
//	att.thangs[position].pre = thang
//
//	att.thangs = append(att.thangs[:position], append([]*Thang{thang}, att.thangs[position:]...)...)
//	return
//}

func main() {

	// Test that Thang is a Thanger
	// https://splice.com/blog/golang-verify-type-implements-interface-compile-time/
	var _ Thanger = (*Thang)(nil)  // success

	all := AllTheThangs{nil}
	t1 := Thang{"T1", nil, nil}
	all.addToEnd(&t1)

	t2 := Thang{"T2", nil, nil}
	all.addToEnd(&t2)
	t4 := Thang{"T4", nil, nil}
	all.addToEnd(&t4)
	t5 := Thang{"T5", nil, nil}
	all.addToEnd(&t5)

	for i, aThang := range all.thangs {
		thang := *aThang
		pre := *thang.getPre()
		post := *thang.getPost()
		fmt.Println(i, thang.getPre(), aThang, thang.getPost())
		fmt.Printf("        %s           %s          %s\n", pre.getName(), thang.getName(), post.getName())
	}

	os.Exit(3	)

	all.printLinkList(false, nil)
	//all.printLinkList(true, nil)
	//
	//t3 := Thang{"T3", nil, nil}
	//all.insert(2, &t3)
	//all.printLinkList(false, nil)
	//all.printLinkList(true, nil)
	//
	//for i, thang := range all.thangs {
	//	//fmt.Println(i, thang.name)
	//	fmt.Printf("%d:              %s <- %s -> %s\n", i, thang.pre.name, thang.name, thang.post.name)
	//	fmt.Printf("%p <- %p -> %p\n", thang.pre, thang, thang.post)

	os.Exit(3)

}