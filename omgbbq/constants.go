package main

// https://splice.com/blog/iota-elegant-constants-golang/

import "fmt"

// https://github.com/golang/go/wiki/Iota
type ByteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

type Suit int

// https://programming.guide/go/iota.html
const (
	Spades   Suit = iota
	Hearts
	Diamonds
	Clubs
)

func (s Suit) String() string {
	return [...]string{"Spades", "Hearts", "Diamonds", "Clubs"}[s]
}

func main() {
	fmt.Println(KB, MB)

	suit1 := Suit(1)

	var suit2 Suit
	suit2 = 2

	var suit3 Suit = Clubs

	fmt.Println(suit1, suit2, suit3)
}
