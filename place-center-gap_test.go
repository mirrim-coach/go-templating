package template

import (
	"testing"
)

func TestPlaceCenterGap0(t *testing.T) {
	ret := placeCenterGap(0, nil)

	if ret != "" {
		t.Errorf("Expected placeCenterGap(0) to return ''. Got: %s", ret)
	}
}
func TestPlaceCenterGap10(t *testing.T) {
	ret := placeCenterGap(10, nil)

	if ret != "left: 50%; width: 8%" {
		t.Errorf("Expected placeCenterGap(10) to return 'left: 50%%; width: 8%%'. Got: %s", ret)
	}
}
func TestPlaceCenterGapNeg10(t *testing.T) {
	ret := placeCenterGap(-10, nil)

	if ret != "left: 8%; width: 8%" {
		t.Errorf("Expected placeCenterGap(-10) to return 'left: 8%%; width: 8%%'. Got: %s", ret)
	}
}
