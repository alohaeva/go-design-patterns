package singleton

import (
	"fmt"
	"sync"
)

// sync.Once ensures lazy initialization and thread safety
var once sync.Once

type Singleton struct{}

// private static instance
var instance *Singleton

// Private constructor equivalent (Go doesn't allow private constructors)
// Instead, we ensure that the instance is only created once.
func newSingleton() *Singleton {
	return &Singleton{}
}

// Get Create method to return the singleton instance
func getSingleton() *Singleton {
	once.Do(func() {
		instance = newSingleton()
	})
	return instance
}

func TestSingleton() {
	//Test the Singleton
	s1 := getSingleton()
	s2 := getSingleton()

	if s1 == s2 {
		fmt.Println("Both are the same instance")
	} else {
		fmt.Println("Different instances")
	}
}
