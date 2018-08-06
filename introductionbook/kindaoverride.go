package main

import "fmt"

// I need to read this many times.
// https://www.ardanlabs.com/blog/2016/10/reducing-type-hierarchies.html

type father struct {
}

// When the son calls this, the f.shout() now refers to father.shout().
// In python and java, the son.shout() would be called.  Because in golang,
// the override works at only the first level?
func (f *father) shoutLouder() {
	fmt.Println("louder...")
	f.shout()
}

func (f father) shout() {
	fmt.Println("father shouting")
}

type son struct {
	father
}

func (s *son) shout() {
	fmt.Println("son shouting")
}

func main() {
	//var s = son{}
	s := son{}
	s.shout()
	s.shoutLouder()
}

//func main() {
//	var s son = son{}
//	s.shout()
//	s.shoutLouder()
//}
///*
//son shouting
//louder...
//father shouting
// */