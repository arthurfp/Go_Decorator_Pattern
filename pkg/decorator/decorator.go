package decorator

// Decorator struct embeds a Component and extends its behavior.
type Decorator struct {
	component Component
}

// NewDecorator creates a new Decorator wrapping the provided Component.
func NewDecorator(component Component) *Decorator {
	return &Decorator{component: component}
}

// Operation returns the result of the wrapped Component's Operation with additional behavior.
func (d *Decorator) Operation() string {
	return "Decorator(" + d.component.Operation() + ")"
}
