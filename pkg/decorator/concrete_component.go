package decorator

// ConcreteComponent is a concrete implementation of the Component interface.
type ConcreteComponent struct{}

func (c *ConcreteComponent) Operation() string {
	return "ConcreteComponent"
}
