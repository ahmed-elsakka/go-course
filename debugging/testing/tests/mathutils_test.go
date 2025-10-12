package mathutils_test

import (
	"testing"

	"testing-example/mathutils"
)

func TestAdd(t *testing.T) {
	result := mathutils.Add(1, 2)
	expected := 3

	if result != expected {
		t.Errorf("Expected result of Add(1, 2) is %d, received value is %d", expected, result)
	}
}

func TestMultiply(t *testing.T) {
	result := mathutils.Multiply(1, 2)
	expected := 2

	if result != expected {
		t.Errorf("Expected result of Multiply(1, 2) is %d, received value is %d", expected, result)
	}
}
