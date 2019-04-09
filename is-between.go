package template

import (
	"github.com/aymerick/raymond"
)

// Handlebars.registerHelper('isBetween', function(value, start, end, options) {
//   return value >= start && value <= end;
// })

func isBetween(input interface{}, low interface{}, high interface{}, options *raymond.Options) bool {
	value, _ := getFloat(input)
	start, _ := getFloat(low)
	end, _ := getFloat(high)
	return value >= start && value <= end
}
