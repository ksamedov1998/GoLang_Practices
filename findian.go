package main

import (
	"fmt"
	"strings"
)

func findian() {
	var someString string
	fmt.Println("Enter a word :")
	fmt.Scan(&someString)
	someString = strings.ToLower(someString)
	switch {
	case strings.HasPrefix(someString, "i") && strings.HasSuffix(someString, "n") && strings.Contains(someString, "a"):
		fmt.Println("Found!")
	default:
		fmt.Println("Not Found!")
	}
	fmt.Scan()
}
