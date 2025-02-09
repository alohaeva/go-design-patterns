package main

import (
	"design-patterns/builder"
	"design-patterns/factory"
	"design-patterns/prototype"
	"design-patterns/singleton"
)

func main() {
	singleton.TestSingleton()
	prototype.TestPrototype()
	factory.TestFactory("email")
	factory.TestFactory("slack")
	factory.TestFactory("sms")
	builder.TestBuilder()
}
