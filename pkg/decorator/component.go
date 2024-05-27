package decorator

// Component interface defines the method that concrete components and decorators must implement.
type Component interface {
	Operation() string
}
