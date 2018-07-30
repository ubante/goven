package main

import "fmt"

/*
Pointers are pointing me;  hard.

I don't understand why after banana.next points to carrot,
apple.next.next still points to nil.
 */

type Item struct {
	name string
	next *Item
}

type Things struct {
	items []*Item
}

func (t *Things) add(name string) {
	newItem := Item{name, nil}
	t.items = append(t.items, &newItem)

	if len(t.items) == 1 {
		return
	}

	t.items[len(t.items)-2].next = &newItem
	fmt.Print(t.items[len(t.items)-2])
	fmt.Print("-->")
	fmt.Println(newItem)
}

func main() {
	var things Things
	things.add("apple")
	things.add("banana")
	things.add("carrot")
	things.add("doughnut")

	fmt.Println("==================")
	for _, i := range things.items {
		fmt.Println(i)
		fmt.Print("-->")
		fmt.Println(i.next)
	}

	fmt.Printf("%s -> %s -> %s -> %s\n", things.items[0].name, things.items[0].next.name,
			things.items[0].next.next.name, things.items[0].next.next.next.name)
}

/*
Output was:

{apple 0xc042002460}-->{banana <nil>}
{banana 0xc0420024e0}-->{carrot <nil>}
{carrot 0xc042002540}-->{doughnut <nil>}
{apple 0xc042002460}
-->&{banana <nil>}
{banana 0xc0420024e0}
-->&{carrot <nil>}
{carrot 0xc042002540}
-->&{doughnut <nil>}
{doughnut <nil>}
--><nil>

Now it is good:
&{apple 0xc042002440}-->{banana <nil>}
&{banana 0xc0420024c0}-->{carrot <nil>}
&{carrot 0xc042002540}-->{doughnut <nil>}
==================
&{apple 0xc042002440}
-->&{banana 0xc0420024c0}
&{banana 0xc0420024c0}
-->&{carrot 0xc042002540}
&{carrot 0xc042002540}
-->&{doughnut <nil>}
&{doughnut <nil>}
--><nil>
apple -> banana -> carrot -> doughnut
 */