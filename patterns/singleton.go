package singleton

import "sync"

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
func Get() *Singleton {
	once.Do(func() {
		instance = newSingleton()
	})
	return instance
}
