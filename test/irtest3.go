package main

import "fmt"

func fact(i int) int {
	if i == 0 {
		return 1
	}
	return i * fact(i-1)
}

func main() {
	var j int = 5
	fact(5)
}