package template

import (
	"github.com/aymerick/raymond"
)

func toFloat(val interface{}, options *raymond.Options) float64 {
	return_value, _ := getFloat(val)
	return return_value
}
