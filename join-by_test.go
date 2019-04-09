package template

import (
	"testing"
)

func TestJoinBy(t *testing.T) {
	list := []string{"a", "b", "c"}

	joined := joinBy(list, ",", nil)

	if joined != "a,b,c" {
		t.Errorf("Expected joinBy(%v, ',') to return 'a,b,c'. Got: %s", list, joined)
	}
}
