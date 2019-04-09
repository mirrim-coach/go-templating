package template

import (
	"testing"
)

func TestNumGroupsOf(t *testing.T) {
	// 5 in groups of 2: round to 3 groups
	ret := numGroupsOf(5, 2, nil)
	if ret != 3 {
		t.Errorf("Expected numGroupsOf(5, 2, nil) to return 3. Got: %f", ret)
	}
}

func TestNumGroupsOfLessThanGroup(t *testing.T) {
	// 5 in groups of 2: round to 3 groups
	ret := numGroupsOf("2", "5", nil)
	if ret != 1 {
		t.Errorf("Expected numGroupsOf(2, 5, nil) to return 1. Got: %f", ret)
	}
}

var numGroupsOfTests = []Test{
	Test{
		name:     "#numGroupsOf",
		input:    "{{numGroupsOf (arrayLength goodbyes) 15}}",
		data:     map[string]interface{}{"goodbyes": []int{1, 2, 3}},
		privData: map[string]interface{}{"exclaim": numGroupsOf},
		helpers:  map[string]interface{}{"numGroupsOf": numGroupsOf, "arrayLength": arrayLength},
		partials: nil,
		output:   "1",
	},
}

func TestNumGroupsOfTests(t *testing.T) {
	launchTests(t, numGroupsOfTests)
}
