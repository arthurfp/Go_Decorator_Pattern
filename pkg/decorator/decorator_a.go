package decorator

// DecoratorA extends the behavior of the component it wraps.
type DecoratorA struct {
	component Component
}

// NewDecoratorA creates a new DecoratorA wrapping the provided Component.
func NewDecoratorA(component Component) *DecoratorA {
	return &DecoratorA{component: component}
}

// Operation returns the result of the wrapped Component's Operation with additional behavior.
func (d *DecoratorA) Operation() string {
	return "DecoratorA(" + d.component.Operation() + ")"
}
