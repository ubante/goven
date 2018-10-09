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
	verbose := false
	fmt.Printf("Current length is: %d, ", len(att.thangs))
	fmt.Println("Adding:", thang.getName())

	if len(att.thangs) == 0 {
		att.thangs = append(att.thangs, &thang)
		return
	}

	if len(att.thangs) == 3 {
		verbose = true
	}

	if verbose {
		fmt.Println("VERBOSE is TRUE")

		fmt.Println("Appending:", thang.getName())
		indexLast := len(att.thangs) - 1
		lastThang := *att.thangs[indexLast]
		lastThang.setPost(&thang)
		thang.setPre(att.thangs[indexLast])

		indexFirst := 0
		firstThang := *att.thangs[indexFirst]
		firstThang.setPre(&thang)
		thang.setPost(att.thangs[indexFirst])

		fmt.Println("The first thang is", att.thangs[indexFirst])
		fmt.Println("The last thang is ", att.thangs[indexLast])

		att.thangs = append(att.thangs, &thang)
		fmt.Println()

		return
	}

	fmt.Println("Appending:", thang.getName())
	indexLast := len(att.thangs) - 1
	lastThang := *att.thangs[indexLast]
	lastThang.setPost(&thang)
	thang.setPre(att.thangs[indexLast])

	indexFirst := 0
	firstThang := *att.thangs[indexFirst]
	firstThang.setPre(&thang)
	thang.setPost(att.thangs[indexFirst])

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

//Position starts counting at 0 and ends at len(att.thangs)-1.
func (att *AllTheThangs) insert(position int, thang Thanger) {
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
		lastThang := *att.thangs[length-1]
		lastThang.setPost(&thang)
		thang.setPre(att.thangs[length-1])

		firstThang := *att.thangs[0]
		firstThang.setPre(&thang)
		thang.setPost(att.thangs[0])

		att.thangs = append([]*Thanger{&thang}, att.thangs...)

		return
	}

	// If the position is at the end of the slice, do this.
	if position == length {
		fmt.Println("Adding to the end.")
		lastThang := *att.thangs[length-1]
		lastThang.setPost(&thang)
		thang.setPre(att.thangs[length-1])

		firstThang := *att.thangs[0]
		firstThang.setPre(&thang)
		thang.setPost(att.thangs[0])

		att.thangs = append(att.thangs, &thang)

		return
	}

	// Otherwise, do this.
	fmt.Println("Adding to position", position)
	middlePre := *att.thangs[position-1]
	middlePre.setPost(&thang)
	thang.setPre(att.thangs[position-1])

	middlePost := *att.thangs[position]
	middlePost.setPre(&thang)
	thang.setPost(att.thangs[position])

	att.thangs = append(att.thangs[:position], append([]*Thanger{&thang}, att.thangs[position:]...)...)

	return
}

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

	//fmt.Println("\n\nJust for fun, let's compare the memory address of a class to the memory of its pointer " +
	//	"dereferenced:")
	//t6 := Thang{"T6", nil, nil}
	//t6p := &t6
	//t6pd := *t6p
	//t6pdp := &t6pd
	//fmt.Printf("%p | %p | %p | %p\n", t6p, t6p, t6pd, t6pdp)
	//os.Exit(3	)

	all.printLinkList(false, nil)
	all.printLinkList(true, nil)

	t3 := Thang{"T3", nil, nil}
	all.insert(2, &t3)
	all.printLinkList(false, nil)
	all.printLinkList(true, nil)

	fmt.Println()
	for i, aThang := range all.thangs {
		thang := *aThang
		pre := *thang.getPre()
		post := *thang.getPost()
		fmt.Println(i, thang.getPre(), aThang, thang.getPost())
		fmt.Printf("        %s           %s          %s\n", pre.getName(), thang.getName(), post.getName())
	}

	// Now add a SubThang.
	t6 := SubThang{Thang{"S6", nil, nil}}
	all.addToEnd(&t6)
	all.printLinkList(false, nil)
	all.printLinkList(true, nil)

	// Add it to the middle
	tx := SubThang{Thang{"SX", nil, nil}}
	all.insert(2, &tx)
	all.printLinkList(false, nil)
	all.printLinkList(true, nil)

}

/*
Current length is: 0, Adding: T1
Current length is: 1, Adding: T2
Appending: T2
Current length is: 2, Adding: T4
Appending: T4
Current length is: 3, Adding: T5
VERBOSE is TRUE
Appending: T5
The first thang is 0xc0420541c0
The last thang is  0xc042054220

0 0xc042054250 0xc0420541c0 0xc0420541e0
        T5           T1          T2
1 0xc0420541c0 0xc0420541e0 0xc042054220
        T1           T2          T4
2 0xc0420541e0 0xc042054220 0xc042054250
        T2           T4          T5
3 0xc042054220 0xc042054250 0xc0420541c0
        T4           T5          T1
T1 -> T2 -> T4 -> T5
T1 <- T5 <- T4 <- T2
Adding to position 2
T1 -> T2 -> T3 -> T4 -> T5
T1 <- T5 <- T4 <- T3 <- T2

0 0xc042054250 0xc0420541c0 0xc0420541e0
        T5           T1          T2
1 0xc0420541c0 0xc0420541e0 0xc042054380
        T1           T2          T3
2 0xc0420541e0 0xc042054380 0xc042054220
        T2           T3          T4
3 0xc042054380 0xc042054220 0xc042054250
        T3           T4          T5
4 0xc042054220 0xc042054250 0xc0420541c0
        T4           T5          T1
Current length is: 5, Adding: S6
Appending: S6
T1 -> T2 -> T3 -> T4 -> T5 -> S6
T1 <- S6 <- T5 <- T4 <- T3 <- T2
Adding to position 2
T1 -> T2 -> SX -> T3 -> T4 -> T5 -> S6
T1 <- S6 <- T5 <- T4 <- T3 <- SX <- T2

Process finished with exit code 0

 */