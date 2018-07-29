package main

import "fmt"

type Cosa struct {
	name string
	next *Cosa
	prev *Cosa
}

func main() {
	a := Cosa{"apple", nil, nil}
	fmt.Println(a)

	b := Cosa{"banana", nil, nil}
	fmt.Println(a, b)
	a.next = &b
	fmt.Println(a, b)

	c := Cosa {"carrot", nil, nil}
	fmt.Println(a, b, c)
	b.next = &c
	fmt.Println(a, b, c)
	fmt.Println(a.next.next)
	fmt.Println(a.next.next.name)
}

/*
This works as expected so my problem is with append() or with how I'm
declaring AnimalQueue.queue as a non-pointer.

{apple <nil> <nil>}
{apple <nil> <nil>} {banana <nil> <nil>}
{apple 0xc042002460 <nil>} {banana <nil> <nil>}
{apple 0xc042002460 <nil>} {banana <nil> <nil>} {carrot <nil> <nil>}
{apple 0xc042002460 <nil>} {banana 0xc042002500 <nil>} {carrot <nil> <nil>}
&{carrot <nil> <nil>}
carrot

 */