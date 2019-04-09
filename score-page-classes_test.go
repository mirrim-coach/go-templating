package template

import (
	"testing"
)

func TestScorePageClasses0(t *testing.T) {
	ret := scorePageClasses(0, nil)

	if ret != " match" {
		t.Errorf("Expected scorePageClasses(0) to return ' match'. Got: %s", ret)
	}
}
func TestScorePageClasses1(t *testing.T) {
	ret := scorePageClasses(1, nil)

	if ret != "" {
		t.Errorf("Expected scorePageClasses(1) to return ''. Got: %s", ret)
	}
}
