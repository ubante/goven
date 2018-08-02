package main

import (
	"os"
	"fmt"
	"strings"
)

// https://gobyexample.com/environment-variables

func main() {
	// To set a key/value pair, use `os.Setenv`. To get a
	// value for a key, use `os.Getenv`. This will return
	// an empty string if the key isn't present in the
	// environment.
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	isEnforced := strings.ToLower(os.Getenv("ETOC"))
	//isEnforced := os.Getenv("ETOC")
	fmt.Println()
	switch isEnforced {
	case "": fmt.Println("Enforcement is not defined; enabling")
	case "False": fmt.Println("We will not enforce.")
	case "false": fmt.Println("We will not enforce.")
	//case "True": fmt.Println("Enabling enforcement.")
	case "true": fmt.Println("Enabling enforcement.")
	default: fmt.Println("Enforcement value is invalid; enabling.")
	}

	// Use `os.Environ` to list all key/value pairs in the
	// environment. This returns a slice of strings in the
	// form `KEY=value`. You can `strings.Split` them to
	// get the key and value. Here we print all the keys.
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}