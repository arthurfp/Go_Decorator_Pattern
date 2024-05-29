package main

import (
	"decorator-pattern-go/pkg/decorator"
	"fmt"
)

func main() {
	component := decorator.ConcreteComponent{}
	fmt.Println("ConcreteComponent Operation:", component.Operation())

	decoratorA := decorator.NewDecoratorA(&component)
	fmt.Println("DecoratorA Operation:", decoratorA.Operation())

	decoratorB := decorator.NewDecoratorB(decoratorA)
	fmt.Println("DecoratorB Operation:", decoratorB.Operation())
}
