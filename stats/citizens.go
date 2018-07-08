package main

// http://spf13.com/post/is-go-object-oriented/

type Citizens struct {
	Country string
	Person
}

type Person struct {
	Name string
	Address Address
}

type Address struct {
	Number string
	Street string
}

func (p *Person) 
func main() {

}