package prototype

import (
	"fmt"
)

// Prototype defines the interface for objects that can be cloned
type Prototype interface {
	Clone() Prototype
	GetName() string
	GetValue() int
}

// ConcretePrototype is a concrete implementation of the Prototype interface
type ConcretePrototype struct {
	Name  string
	Value int
}

// Clone creates a copy of ConcretePrototype1
func (p *ConcretePrototype) Clone() Prototype {
	return &ConcretePrototype{
		Name:  p.Name,
		Value: p.Value,
	}
}

func (p *ConcretePrototype) GetName() string {
	return p.Name
}

func (p *ConcretePrototype) GetValue() int {
	return p.Value
}

// TestPrototype Client function to demonstrate cloning
func TestPrototype() {
	// Creating initial prototypes
	prototype1 := &ConcretePrototype{Name: "Prototype1", Value: 42}

	// Cloning prototypes
	clone1 := prototype1.Clone()

	// Displaying original and cloned objects
	fmt.Printf("Original: %s, Value: %d\n", prototype1.GetName(), prototype1.Value)
	fmt.Printf("Cloned: %s, Value: %d\n", clone1.GetName(), clone1.GetValue())
}
