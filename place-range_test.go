package template

import (
	"testing"
)

func TestPlaceRange(t *testing.T) {
	target_range := map[string]interface{}{"start": -18, "end": 34}
	ret := placeRange(target_range, nil)

	if ret != "width: 305.760000pt; height: 305.760000pt; border-radius: 50%; left: 30.000000%; top: -139.880000pt;" {
		t.Errorf("Expected placeScore(0) to return width: 305.760000pt; height: 305.760000pt; border-radius: 50%%; left: 30.000000%%; top: -139.880000pt;. Got: %s", ret)
	}
}
