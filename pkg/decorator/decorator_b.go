package decorator

// DecoratorB extends the behavior of the component it wraps.
type DecoratorB struct {
	component Component
}

// NewDecoratorB creates a new DecoratorB wrapping the provided Component.
func NewDecoratorB(component Component) *DecoratorB {
	return &DecoratorB{component: component}
}

// Operation returns the result of the wrapped Component's Operation with additional behavior.
func (d *DecoratorB) Operation() string {
	return "DecoratorB(" + d.component.Operation() + ")"
}
