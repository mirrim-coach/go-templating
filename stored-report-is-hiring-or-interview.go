package template

import (
	"github.com/aymerick/raymond"
)

func storedReportIsHiringOrInterview(stored map[string]interface{}, options *raymond.Options) bool {
	if _, ok := stored["type"]; !ok {
		return false
	}

	return stored["type"] == "Hiring" || stored["type"] == "Interview"
}
