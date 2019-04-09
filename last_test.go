package template

import (
	"testing"
)

func TestLast(t *testing.T) {
	list := []string{"a", "b", "c"}

	last := last(list, nil)

	if last != "c" {
		t.Errorf("Expected last([]string{a b c}) to return 'c'. Got: %v", last)
	}
}
