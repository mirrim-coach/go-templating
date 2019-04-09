package template

import (
	"math"
	"testing"
)

func TestToFloatString0(t *testing.T) {
	ret := toFloat("0", nil)

	if ret != 0 {
		t.Errorf("Expected toFloat(\"0\", nil) to return zero. Got: %f", ret)
	}
}

func TestToFloatInt0(t *testing.T) {
	ret := toFloat(0, nil)

	if ret != 0 {
		t.Errorf("Expected toFloat(0, nil) to return zero. Got: %f", ret)
	}
}

func TestToFloatNil(t *testing.T) {
	ret := toFloat(nil, nil)

	if math.IsNaN(ret) == false {
		t.Errorf("Expected toFloat(nil, nil) to return NaN. Got: %f", ret)
	}
}
