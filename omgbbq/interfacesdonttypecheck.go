package main

import "fmt"

/*
https://medium.com/@saiyerram/go-interfaces-pointers-4d1d98d5c9c6
"It seems the function mutator defies the general go-lang type checking.
For e.g. this would fail with normal types, but works with the interfaces."

This is a insane.
*/

//(1) The interface
type Mutable interface {
	mutate(newValue string) error
}
//(2) Struct
type Data struct {
	name string
}
//(3) Implements the interface with a pointer receiver
func (d *Data) mutate(newValue string) error {
	d.name = newValue
	return nil
}
//(4) Function that accepts the interface
func mutator(mute Mutable) error {
	return mute.mutate("mutate")
}
func main() {
	d := Data{name: "fresh"}
	fmt.Println(d.name) //fresh
	//(5) pass as a pointer
	mutator(&d)
	fmt.Println(d.name) //mutate


}

