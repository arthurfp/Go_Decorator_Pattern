package decorator

import "testing"

func TestConcreteComponent_Operation(t *testing.T) {
	component := &ConcreteComponent{}
	expected := "ConcreteComponent"
	result := component.Operation()

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
