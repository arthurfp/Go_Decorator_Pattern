package decorator

import "testing"

func TestDecoratorB_Operation(t *testing.T) {
	component := &ConcreteComponent{}
	decoratorA := NewDecoratorA(component)
	decoratorB := NewDecoratorB(decoratorA)
	expected := "DecoratorB(DecoratorA(ConcreteComponent))"
	result := decoratorB.Operation()

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
