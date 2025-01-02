package main

import "fmt"

var version string

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

func main() {
	fmt.Printf("Version: %s\n", version)
}
