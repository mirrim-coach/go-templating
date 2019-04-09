package template

import (
	"testing"
)

func TestKeys(t *testing.T) {
	list := map[string]interface{}{"a": true, "b": false, "c": true}

	keys := keys(list, nil)

	if !uniqueArray(keys, []string{"a", "b", "c"}) {
		t.Errorf("Expected %v to return keys [a b c]. Got: %v", list, keys)
	}
}
