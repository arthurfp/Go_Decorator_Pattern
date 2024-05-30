package decorator

import "testing"

func TestDecoratorA_Operation(t *testing.T) {
	component := &ConcreteComponent{}
	decoratorA := NewDecoratorA(component)
	expected := "DecoratorA(ConcreteComponent)"
	result := decoratorA.Operation()

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
