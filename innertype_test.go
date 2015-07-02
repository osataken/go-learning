package learning

import (
	"testing"
)

func TestInnerTypeFunctionShouldBeCall(t *testing.T) {
	outer := Outer{
		Inner : Inner {},
	}
	expected := "inner method"

	result := outer.Inner.doSomething()

	if result != expected {
		t.Errorf("Expected %V but was %V", expected, result)
	}
}

func TestOuterTypeFunctionShouldBeCall(t *testing.T) {
	outer := Outer{
		Inner : Inner {},
	}
	expected := "outer method"

	result := outer.doSomething()

	if result != expected {
		t.Errorf("Expected %V but was %V", expected, result)
	}
}
