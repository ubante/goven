package main

import (
	"time"
	"fmt"
)

// https://stackoverflow.com/questions/24122821/go-golang-time-now-unixnano-convert-to-milliseconds

func main() {
	for i := 1; i < 10; i++ {
		a := makeTimestamp()
		fmt.Printf("%d \n", a)
	}
}

func makeTimestamp() int64 {
	return time.Now().UnixNano()
	//return time.Now().UnixNano() / int64(time.Millisecond)
}
