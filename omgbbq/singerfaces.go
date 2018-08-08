package main

import "fmt"

type Singer interface {
	start()
	stop()
}

type GenericSinger struct {
	name string
	octave int
}

func (gs GenericSinger) start() {
	fmt.Printf("%s starts to sing.\n", gs.name)
}

func (gs GenericSinger) stop() {
	fmt.Printf("%s stops singing.\n", gs.name)
}

type SopranoSinger struct {
	GenericSinger
}

func (ss SopranoSinger) stop() {
	fmt.Println("Sopranos never stop singing.")
}

type TenorSinger struct {
	GenericSinger
}

func (ts TenorSinger) start() {
	fmt.Printf("%s starts to sing and the crowd starts to tear.\n", ts.name)
}

func main() {
	lea := SopranoSinger{GenericSinger{name: "Lea", octave: 4}}
	andrea := TenorSinger{GenericSinger{name: "Andrea", octave: 2}}
	theCond := GenericSinger{name: "Bun", octave: 0}

	singers := []Singer{lea, andrea, theCond}

	for _, singer := range singers {
		singer.start()
		defer singer.stop()
	}
}
