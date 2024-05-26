package decorator

// Component interface defines the method that concrete components and decorators must implement.
type Component interface {
	Operation() string
}

// ConcreteComponent is a concrete implementation of the Component interface.
type ConcreteComponent struct{}

func (c *ConcreteComponent) Operation() string {
	return "ConcreteComponent"
}
