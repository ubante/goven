package other

import (
	"fmt"
	//"goven/poker/main"
	"goven/omgbbq/yetAnother"
)

type SomeOtherThing struct {
	innards string
	InnardsCapitalized string
}

func (sot SomeOtherThing) CallBack() {
	fmt.Println("I'm in S-O-T.")

	var sdt yetAnother.SisterDirThing
	sdt.Guts = "sisterguts"
	fmt.Println("This is from a type in a sister dir that is a subdir of the main dir:", sdt)

	aot := AnotherOtherThing{ootString: "OOT"}
	fmt.Println("this is the other thing in other oustide of this:", aot)

}

