package template

import (
	"testing"
)

func TestCategoryIsHiringEmpty(t *testing.T) {
	input := map[string]interface{}{}

	ret := categoryIsHiring(input, nil)

	if ret {
		t.Errorf("Expected categoryIsHiring(%v) to return false. Got: %t", input, ret)
	}
}
func TestCategoryIsHiringNone(t *testing.T) {
	input := map[string]interface{}{"name": "none"}

	ret := categoryIsHiring(input, nil)

	if ret {
		t.Errorf("Expected categoryIsHiring(%v) to return false. Got: %t", input, ret)
	}
}
func TestCategoryIsHiringHiring(t *testing.T) {
	input := map[string]interface{}{"name": "Hiring"}

	ret := categoryIsHiring(input, nil)

	if !ret {
		t.Errorf("Expected categoryIsHiring(%v) to return true. Got: %t", input, ret)
	}
}
