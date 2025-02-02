package main

import (
	"design-patterns/patterns"
	"fmt"
)

func main() {
	// Test the Singleton
	s1 := singleton.Get()
	s2 := singleton.Get()

	if s1 == s2 {
		fmt.Println("Both are the same instance")
	} else {
		fmt.Println("Different instances")
	}
}
