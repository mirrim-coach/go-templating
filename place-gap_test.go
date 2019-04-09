package template

import (
	"testing"
)

func TestPlaceGap(t *testing.T) {
	r := map[string]interface{}{"start": 5, "end": 6}

	ret := placeGap(5, 10, r, nil)

	if ret != "left: 54.750000%; width: -1.000000%;" {
		t.Errorf("Expected placeScore(0) to return 'left: 54.750000%%; width: -1.000000%%;'. Got: %s", ret)
	}
}
