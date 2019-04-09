package template

import (
	"testing"
)

func TestIsBetween0_10(t *testing.T) {
	ret := isBetween(3, 0, 10, nil)
	if ret != true {
		t.Error("Expected isBetween(3, 0, 10, nil) to return true. Got: false")
	}
}

func TestIsBetween5_10(t *testing.T) {
	ret := isBetween(3, 5, 10, nil)
	if ret != false {
		t.Error("Expected isBetween(3, 5, 10, nil) to return false. Got: true")
	}
}
