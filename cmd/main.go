package main

import (
	"decorator-pattern-go/pkg/decorator"
	"fmt"
)

func main() {
	component := decorator.ConcreteComponent{}
	fmt.Println(component.Operation())
}
