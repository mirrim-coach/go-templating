package template

import (
	"github.com/aymerick/raymond"
)

func scorePageClasses(input interface{}, options *raymond.Options) string {
	gap, _ := getFloat(input)

	if gap == 0 {
		return " match"
	}
	return ""
}
