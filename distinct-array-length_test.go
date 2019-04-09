package template

import (
	// log "github.com/sirupsen/logrus"
	"testing"
)

func TestDistinctArrayLengthStringArray(t *testing.T) {
	list := []string{"1", "1", "2", "3", "4", "4", "5", "6"}

	length := distinctArrayLength(list, nil)

	if length != 6 {
		t.Errorf("Expected slice %v to have a unique length of 6. Got: %d", list, length)
	}
}

func TestDistinctArrayLengthNumberArray(t *testing.T) {
	list := []int{1, 1, 2, 3, 4, 4, 5, 6}

	length := distinctArrayLength(list, nil)

	if length != 6 {
		t.Errorf("Expected slice %v to have a unique length of 6. Got: %d", list, length)
	}
}
