package main

import (
	"fmt"
	"gopkg.in/inconshreveable/log15.v2"
)

func main() {
	fmt.Println("Testing the logs.")
	log15.Info("Infof doesn't work on desktop... hrm")
}
