package main

import (
	"fmt"
	"regexp"
	"strings"
	"unsafe"
)

// http://spf13.com/post/is-go-object-oriented/

type Segment string

type Point struct {
	Val   float64
	Empty bool
}

type Data struct {
	// Map from segment key to list of data points.
	series               map[Segment][]Point
	start, end, interval int64
}

func (k Segment) Explode() map[string]string {
	tags := map[string]string{}
	for _, pair := range strings.Split(string(k), " ") {
		words := strings.Split(pair, "=")
		if len(words) == 2 {
			tags[words[0]] = words[1]
		}
	}
	return tags
}

//func (p *Person)
func main() {
	//var name = "dsl"
	name := "dsl"
	fmt.Printf("Name is %s", name)
	var patterns [2]string
	patterns[0] = "ad_video_view_hf_lr_daily_scorpion"
	patterns[1] = "test_model"

	var exprs []*regexp.Regexp
	for _, pattern := range patterns {
		fmt.Println(pattern)
		expr, err := regexp.Compile(pattern)
		if err == nil {
			exprs = append(exprs, expr)
		}
	}

	//var d Data
	//for segment := range d.series {
	//	tags := segment.Explode()
	//	var val float64

	//for _, expr := range exprs {
	//	if expr.MatchString(ta)
	//}

	a := [2]int{1, 2}
	fmt.Println(a)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	i := 111
	//var pi *int
	//pi := *i
	vi := &i

	fmt.Printf("\ni:        %v\nmem_addr: %v\n", i, vi)
	//fmt.Printf("i:   %v\nptr: %v\nmem_addr: %v\n", i, pi, vi)

	str1 := "abc"
	//str1 := Segment("abc")
	str1copy := str1
	str1ptr1 := &str1
	str1ptr2 := &str1
	str1ptr11 := &str1ptr1
	//fmt.Printf("String = %s; Copy = %s\n", str1, str1copy)
	fmt.Printf("String = %s; Copy = %s; Pointer = %s; Pointer = %s\n",
		str1, str1copy, str1ptr1, str1ptr2)
	fmt.Printf("pointer = %s; deref = %s\n", str1ptr1, str1)
	fmt.Printf("ptrptr = %s; pointer = %s, deref = %s\n",
		str1ptr11, *str1ptr11, **str1ptr11)
	//fmt.Println(str1ptr11)

	str1 = "xyz"
	fmt.Println("\nAfter switch str1:")
	//fmt.Printf("String = %s; Copy = %s\n", str1, str1copy)
	fmt.Printf("String = %s; Copy = %s; Pointer = %s\n",
		str1, str1copy, str1ptr1)

	fmt.Println(unsafe.Sizeof(str1))
	fmt.Println(unsafe.Sizeof(str1ptr1))
	fmt.Println(unsafe.Sizeof(str1ptr11))
}
