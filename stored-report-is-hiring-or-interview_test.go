package template

import (
	"testing"
)

func TestStoredReportIsHiringOrInterviewEmpty(t *testing.T) {
	input := map[string]interface{}{}

	ret := storedReportIsHiringOrInterview(input, nil)

	if ret {
		t.Errorf("Expected storedReportIsHiringOrInterview(%v) to return false. Got: %t", input, ret)
	}
}

func TestStoredReportIsHiringOrInterviewNone(t *testing.T) {
	input := map[string]interface{}{"type": "none"}

	ret := storedReportIsHiringOrInterview(input, nil)

	if ret {
		t.Errorf("Expected storedReportIsHiringOrInterview(%v) to return false. Got: %t", input, ret)
	}
}

func TestStoredReportIsHiringOrInterviewHiring(t *testing.T) {
	input := map[string]interface{}{"type": "Hiring"}

	ret := storedReportIsHiringOrInterview(input, nil)

	if !ret {
		t.Errorf("Expected storedReportIsHiringOrInterview(%v) to return true. Got: %t", input, ret)
	}
}
func TestStoredReportIsHiringOrInterviewInterview(t *testing.T) {
	input := map[string]interface{}{"type": "Interview"}

	ret := storedReportIsHiringOrInterview(input, nil)

	if !ret {
		t.Errorf("Expected storedReportIsHiringOrInterview(%v) to return true. Got: %t", input, ret)
	}
}
