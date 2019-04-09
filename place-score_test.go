package template

import (
	"testing"
)

func TestPlaceScore0(t *testing.T) {
	ret := placeScore(0, nil)

	if ret != "zero" {
		t.Errorf("Expected placeScore(0) to return zero. Got: %s", ret)
	}
}
func TestPlaceScore1(t *testing.T) {
	ret := placeScore(1, nil)

	if ret != "positive score-1" {
		t.Errorf("Expected placeScore(1) to return 'positive score-1'. Got: %s", ret)
	}
}
func TestPlaceScoreNeg1(t *testing.T) {
	ret := placeScore(-1, nil)

	if ret != "negative score-1" {
		t.Errorf("Expected placeScore(1) to return 'negative score-1'. Got: %s", ret)
	}
}
