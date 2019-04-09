package template

import (
	"github.com/aymerick/raymond"
	"strings"
)

func categoryIsHiring(category map[string]interface{}, options *raymond.Options) bool {
	if _, ok := category["name"]; !ok {
		return false
	}

	return strings.ToLower(category["name"].(string)) == "hiring"
}
