package main

import (
	"decorator-pattern-go/pkg/decorator"
	"fmt"
)

func main() {
	component := decorator.ConcreteComponent{}
	fmt.Println("ConcreteComponent Operation:", component.Operation())

	decorator := decorator.NewDecorator(&component)
	fmt.Println("Decorator Operation:", decorator.Operation())
}
